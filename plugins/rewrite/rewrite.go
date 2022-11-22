package plugins

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Rewriter struct {
	pkg    *packages.Package
	files  map[string]string
	copied map[ast.Decl]bool
}

func New(dir string) (*Rewriter, error) {
	importPath := ImportPathForDir(dir)
	if importPath == "" {
		return nil, fmt.Errorf("import path not found for directory: %q", dir)
	}
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes,
	}, importPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("package not found for importPath: %s", importPath)
	}

	return &Rewriter{
		pkg:    pkgs[0],
		files:  map[string]string{},
		copied: map[ast.Decl]bool{},
	}, nil
}

func (r *Rewriter) getSource(start, end token.Pos) string {
	startPos := r.pkg.Fset.Position(start)
	endPos := r.pkg.Fset.Position(end)

	if startPos.Filename != endPos.Filename {
		panic("cant get source spanning multiple files")
	}

	file := r.getFile(startPos.Filename)
	return file[startPos.Offset:endPos.Offset]
}

func (r *Rewriter) getFile(filename string) string {
	if _, ok := r.files[filename]; !ok {
		b, err := os.ReadFile(filename)
		if err != nil {
			panic(fmt.Errorf("unable to load file, already exists: %w", err))
		}

		r.files[filename] = string(b)

	}

	return r.files[filename]
}

func (r *Rewriter) GetMethodComment(structname string, methodname string) string {
	for _, f := range r.pkg.Syntax {
		for _, d := range f.Decls {
			d, isFunc := d.(*ast.FuncDecl)
			if !isFunc {
				continue
			}
			if d.Name.Name != methodname {
				continue
			}
			if d.Recv == nil || len(d.Recv.List) == 0 {
				continue
			}
			recv := d.Recv.List[0].Type
			if star, isStar := recv.(*ast.StarExpr); isStar {
				recv = star.X
			}
			ident, ok := recv.(*ast.Ident)
			if !ok {
				continue
			}

			if ident.Name != structname {
				continue
			}
			return d.Doc.Text()
		}
	}

	return ""
}
func (r *Rewriter) GetMethodBody(structname string, methodname string) string {
	for _, f := range r.pkg.Syntax {
		for _, d := range f.Decls {
			d, isFunc := d.(*ast.FuncDecl)
			if !isFunc {
				continue
			}
			if d.Name.Name != methodname {
				continue
			}
			if d.Recv == nil || len(d.Recv.List) == 0 {
				continue
			}
			recv := d.Recv.List[0].Type
			if star, isStar := recv.(*ast.StarExpr); isStar {
				recv = star.X
			}
			ident, ok := recv.(*ast.Ident)
			if !ok {
				continue
			}

			if ident.Name != structname {
				continue
			}

			r.copied[d] = true

			return r.getSource(d.Body.Pos()+1, d.Body.End()-1)
		}
	}

	return ""
}

func (r *Rewriter) MarkStructCopied(name string) {
	for _, f := range r.pkg.Syntax {
		for _, d := range f.Decls {
			d, isGen := d.(*ast.GenDecl)
			if !isGen {
				continue
			}
			if d.Tok != token.TYPE || len(d.Specs) == 0 {
				continue
			}

			spec, isTypeSpec := d.Specs[0].(*ast.TypeSpec)
			if !isTypeSpec {
				continue
			}

			if spec.Name.Name != name {
				continue
			}

			r.copied[d] = true
		}
	}
}

func (r *Rewriter) ExistingImports(filename string) []Import {
	filename, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	for _, f := range r.pkg.Syntax {
		pos := r.pkg.Fset.Position(f.Pos())

		if filename != pos.Filename {
			continue
		}

		var imps []Import
		for _, i := range f.Imports {
			name := ""
			if i.Name != nil {
				name = i.Name.Name
			}
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				panic(err)
			}
			imps = append(imps, Import{name, path})
		}
		return imps
	}
	return nil
}

func (r *Rewriter) RemainingSource(filename string) string {
	filename, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	for _, f := range r.pkg.Syntax {
		pos := r.pkg.Fset.Position(f.Pos())

		if filename != pos.Filename {
			continue
		}

		var buf bytes.Buffer

		for _, d := range f.Decls {
			if r.copied[d] {
				continue
			}

			if d, isGen := d.(*ast.GenDecl); isGen && d.Tok == token.IMPORT {
				continue
			}

			buf.WriteString(r.getSource(d.Pos(), d.End()))
			buf.WriteString("\n")
		}

		return strings.TrimSpace(buf.String())
	}
	return ""
}

type Import struct {
	Alias      string
	ImportPath string
}

var gopaths []string

func init() {
	gopaths = filepath.SplitList(build.Default.GOPATH)
	for i, p := range gopaths {
		gopaths[i] = filepath.ToSlash(filepath.Join(p, "src"))
	}
}

// NameForDir manually looks for package stanzas in files located in the given directory. This can be
// much faster than having to consult go list, because we already know exactly where to look.
func NameForDir(dir string) string {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return SanitizePackageName(filepath.Base(dir))
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return SanitizePackageName(filepath.Base(dir))
	}
	fset := token.NewFileSet()
	for _, file := range files {
		if !strings.HasSuffix(strings.ToLower(file.Name()), ".go") {
			continue
		}

		filename := filepath.Join(dir, file.Name())
		if src, err := parser.ParseFile(fset, filename, nil, parser.PackageClauseOnly); err == nil {
			return src.Name.Name
		}
	}

	return SanitizePackageName(filepath.Base(dir))
}

type goModuleSearchResult struct {
	path       string
	goModPath  string
	moduleName string
}

var goModuleRootCache = map[string]goModuleSearchResult{}

// goModuleRoot returns the root of the current go module if there is a go.mod file in the directory tree
// If not, it returns false
func goModuleRoot(dir string) (string, bool) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	dir = filepath.ToSlash(dir)

	dirs := []string{dir}
	result := goModuleSearchResult{}

	for {
		modDir := dirs[len(dirs)-1]

		if val, ok := goModuleRootCache[dir]; ok {
			result = val
			break
		}

		if content, err := os.ReadFile(filepath.Join(modDir, "go.mod")); err == nil {
			moduleName := extractModuleName(content)
			result = goModuleSearchResult{
				path:       moduleName,
				goModPath:  modDir,
				moduleName: moduleName,
			}
			goModuleRootCache[modDir] = result
			break
		}

		if modDir == "" || modDir == "." || modDir == "/" || strings.HasSuffix(modDir, "\\") {
			// Reached the top of the file tree which means go.mod file is not found
			// Set root folder with a sentinel cache value
			goModuleRootCache[modDir] = result
			break
		}

		dirs = append(dirs, filepath.Dir(modDir))
	}

	// create a cache for each path in a tree traversed, except the top one as it is already cached
	for _, d := range dirs[:len(dirs)-1] {
		if result.moduleName == "" {
			// go.mod is not found in the tree, so the same sentinel value fits all the directories in a tree
			goModuleRootCache[d] = result
		} else {
			if relPath, err := filepath.Rel(result.goModPath, d); err != nil {
				panic(err)
			} else {
				path := result.moduleName
				relPath := filepath.ToSlash(relPath)
				if !strings.HasSuffix(relPath, "/") {
					path += "/"
				}
				path += relPath

				goModuleRootCache[d] = goModuleSearchResult{
					path:       path,
					goModPath:  result.goModPath,
					moduleName: result.moduleName,
				}
			}
		}
	}

	res := goModuleRootCache[dir]
	if res.moduleName == "" {
		return "", false
	}
	return res.path, true
}

func extractModuleName(content []byte) string {
	for {
		advance, tkn, err := bufio.ScanLines(content, false)
		if err != nil {
			panic(fmt.Errorf("error parsing mod file: %w", err))
		}
		if advance == 0 {
			break
		}
		s := strings.Trim(string(tkn), " \t")
		if len(s) != 0 && !strings.HasPrefix(s, "//") {
			break
		}
		if advance <= len(content) {
			content = content[advance:]
		}
	}
	moduleName := string(modregex.FindSubmatch(content)[1])
	return moduleName
}

// ImportPathForDir takes a path and returns a golang import path for the package
func ImportPathForDir(dir string) (res string) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	dir = filepath.ToSlash(dir)

	modDir, ok := goModuleRoot(dir)
	if ok {
		return modDir
	}

	for _, gopath := range gopaths {
		if len(gopath) < len(dir) && strings.EqualFold(gopath, dir[0:len(gopath)]) {
			return dir[len(gopath)+1:]
		}
	}

	return ""
}

var modregex = regexp.MustCompile(`module ([^\s]*)`)

var invalidPackageNameChar = regexp.MustCompile(`[^\w]`)

func SanitizePackageName(pkg string) string {
	return invalidPackageNameChar.ReplaceAllLiteralString(filepath.Base(pkg), "_")
}