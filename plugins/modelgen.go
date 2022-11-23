package plugins

import (
	_ "embed"
	"fmt"
	"go/types"
	"log"
	"sort"
	"strings"
	"text/template"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
)

//go:embed modelgen.gotpl
var modelTemplate string

type BuildMutateHook = func(b *ModelBuild) *ModelBuild

func defaultBuildMutateHook(b *ModelBuild) *ModelBuild {
	return b
}

type ModelBuild struct {
	PackageName string
	Interfaces  []*Interface
	Models      []*Object
	Enums       []*Enum
	Scalars     []string
}

type Interface struct {
	Description string
	Name        string
}

type Object struct {
	Description string
	Name        string
	Fields      []*Field
	Implements  []string
}

type Field struct {
	Description string
	Name        string
	Type        types.Type
	Tag         string
	Gorm        string
}

type Enum struct {
	Description string
	Name        string
	Values      []*EnumValue
}

type EnumValue struct {
	Description string
	Name        string
}

type Plugin struct {
	MutateHook BuildMutateHook
}

func New() plugin.Plugin {
	return &Plugin{
		MutateHook: defaultBuildMutateHook,
	}
}

var _ plugin.ConfigMutator = &Plugin{}

func (m *Plugin) Name() string {
	return "modelgen"
}

var order = make(map[string]struct{}, 0)

func (m *Plugin) MutateConfig(cfg *config.Config) error {

	log.Println("working in model")

	binder := cfg.NewBinder()

	b := &ModelBuild{
		PackageName: cfg.Model.Package,
	}

	for _, schemaType := range cfg.Schema.Types {

		if schemaType.BuiltIn {
			continue
		}

		switch schemaType.Kind {
		case ast.Interface, ast.Union:
			it := &Interface{
				Description: schemaType.Description,
				Name:        schemaType.Name,
			}

			b.Interfaces = append(b.Interfaces, it)
		case ast.Object, ast.InputObject:
			if schemaType == cfg.Schema.Query || schemaType == cfg.Schema.Mutation || schemaType == cfg.Schema.Subscription {
				continue
			}
			it := &Object{
				Description: schemaType.Description,
				Name:        schemaType.Name,
			}
			for _, implementor := range cfg.Schema.GetImplements(schemaType) {
				it.Implements = append(it.Implements, implementor.Name)
			}

			fieldtypes := make([]types.Type, 0)

			for _, field := range schemaType.Fields {
				var typ types.Type

				fieldDef := cfg.Schema.Types[field.Type.Name()]

				if cfg.Models.UserDefined(field.Type.Name()) {
					var err error
					typ, err = binder.FindTypeFromName(cfg.Models[field.Type.Name()].Model[0])
					if err != nil {
						return err
					}
				} else {
					switch fieldDef.Kind {
					case ast.Scalar:
						// no user defined model, referencing a default scalar
						typ = types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), "string", nil),
							nil,
							nil,
						)

					case ast.Interface, ast.Union:
						// no user defined model, referencing a generated interface type
						typ = types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
							types.NewInterfaceType([]*types.Func{}, []types.Type{}),
							nil,
						)

					case ast.Enum:
						// no user defined model, must reference a generated enum
						typ = types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
							nil,
							nil,
						)

					case ast.Object, ast.InputObject:
						// no user defined model, must reference a generated struct
						typ = types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
							types.NewStruct(nil, nil),
							nil,
						)

					default:
						panic(fmt.Errorf("unknown ast type %s", fieldDef.Kind))
					}
				}

				name := field.Name
				if nameOveride := cfg.Models[schemaType.Name].Fields[field.Name].FieldName; nameOveride != "" {
					name = nameOveride
				}

				typ = binder.CopyModifiersFromAst(field.Type, typ)

				gormType := ""

				if isStruct(typ) && (fieldDef.Kind == ast.Object || fieldDef.Kind == ast.InputObject) {
					typ = types.NewPointer(typ)

					it.Fields = append(it.Fields, &Field{
						Name: name + "ID",
						Type: types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), "string", nil),
							nil,
							nil,
						),
						Description: field.Description,
						Tag:         `json:"` + strings.ToLower(field.Name) + `_id"`,
						Gorm:        `gorm:"unique"`,
					})

					gormType = fmt.Sprintf(`gorm:"foreignKey:%s"`, strings.Title(name)+"ID")


				}

				if isArray(typ) && (fieldDef.Kind == ast.Object || fieldDef.Kind == ast.InputObject) {

					typ.String()
					field := &Field{
						Name: name + "ID",
						Type: types.NewNamed(
							types.NewTypeName(0, cfg.Model.Pkg(), "string", nil),
							nil,
							nil,
						),
						Description: field.Description,
						Tag:         `json:"` + strings.ToLower(field.Name) + `_id"`,
						Gorm:        `gorm:"unique"`,
					}
					it.Fields = append(it.Fields, field)

					gormType = fmt.Sprintf(`gorm:"foreignKey:ID;references:%s"`, strings.Title(name)+"ID")


				}

				/*directive := field.Directives.ForName("mapping")
				if directive != nil {
					arg := directive.Arguments.ForName("type")
					if arg != nil {
						if arg.Value.Raw == "many2many" {
							gormType = fmt.Sprintf(`gorm:"many2many:%s;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`, name+"_"+field.Type.Name())

						} else if arg.Value.Raw == "one2many" {
							gormType = fmt.Sprintf(`gorm:"foriegnKey:%s;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`, name)

						} else if arg.Value.Raw == "unique" {
							gormType = fmt.Sprintf(`gorm:"column:%s;uniqueIndex"`, field.Name)

						}else if arg.Value.Raw == "primarykey" {
							gormType = fmt.Sprintf(`gorm:"autoIncrement; primaryKey"`)

						}else if arg.Value.Raw == "autoincrement" {
							gormType = fmt.Sprintf(`gorm:"autoIncrement"`)

						}else if arg.Value.Raw == "notnull" {
							gormType = fmt.Sprintf(`gorm:"not null"`)

						}else if arg.Value.Raw == "index" {
							gormType = fmt.Sprintf(`gorm:"index"`)

						}

					} else {
						gormType = fmt.Sprintf(`gorm:"column:%s"`, field.Name)
					}
				}*/
				value:=getTagFromDirectives(field,schemaType.Name)
				if value!=""{
				gormType = value
				}

				it.Fields = append(it.Fields, &Field{
					Name:        name,
					Type:        typ,
					Description: field.Description,
					Tag:         `json:"` + field.Name + `"`,
					Gorm:        gormType,
				})
				fieldtypes = append(fieldtypes, typ)
			}

			if _, ok := order[it.Name]; !ok {
				order[it.Name] = struct{}{}
			}

			b.Models = append(b.Models, it)
		case ast.Enum:
			it := &Enum{
				Name:        schemaType.Name,
				Description: schemaType.Description,
			}

			for _, v := range schemaType.EnumValues {
				it.Values = append(it.Values, &EnumValue{
					Name:        v.Name,
					Description: v.Description,
				})
			}

			b.Enums = append(b.Enums, it)
		case ast.Scalar:
			b.Scalars = append(b.Scalars, schemaType.Name)
		}
	}
	sort.Slice(b.Enums, func(i, j int) bool { return b.Enums[i].Name < b.Enums[j].Name })
	sort.Slice(b.Models, func(i, j int) bool { return b.Models[i].Name < b.Models[j].Name })
	sort.Slice(b.Interfaces, func(i, j int) bool { return b.Interfaces[i].Name < b.Interfaces[j].Name })

	for _, it := range b.Enums {
		cfg.Models.Add(it.Name, cfg.Model.ImportPath()+"."+templates.ToGo(it.Name))
	}
	for _, it := range b.Models {
		cfg.Models.Add(it.Name, cfg.Model.ImportPath()+"."+templates.ToGo(it.Name))
	}
	for _, it := range b.Interfaces {
		cfg.Models.Add(it.Name, cfg.Model.ImportPath()+"."+templates.ToGo(it.Name))
	}
	for _, it := range b.Scalars {
		cfg.Models.Add(it, "github.com/99designs/gqlgen/graphql.String")
	}

	if len(b.Models) == 0 && len(b.Enums) == 0 && len(b.Interfaces) == 0 && len(b.Scalars) == 0 {
		return nil
	}

	return templates.Render(templates.Options{
		PackageName:     cfg.Model.Package,
		Filename:        cfg.Model.Filename,
		Data:            b,
		GeneratedHeader: true,
		Funcs: template.FuncMap{
			"SetMapValue": SetMapValue,
			"concat":      concatString,
		},
		Packages: cfg.Packages,
		Template: modelTemplate,
	})
}

func isStruct(t types.Type) bool {
	_, is := t.Underlying().(*types.Struct)
	return is
}

func isArray(t types.Type) bool {
	_, is := t.Underlying().(*types.Slice)
	if is {
		_, k := t.Underlying().(*types.Slice).Underlying().(*types.Struct)
		log.Println(k)
	}
	return is
}



var structs = make(map[string]interface{})

func SetMapValue(key string, value string) string {
	structs[key] = "&model." + value
	log.Panicln(structs)
	return ""

}

func concatString(key string) string {
	log.Panicln(key)
	return "$model." + key

}

func getTagFromDirectives(field *ast.FieldDefinition,schemaName string) string {

	gormType := ""

	directive := field.Directives.ForName("mapping")
	if directive != nil {

		arg := directive.Arguments.ForName("type")

		if arg != nil {

			if arg.Value.Raw == "many2many" {

				gormType = fmt.Sprintf(`gorm:"many2many:%s;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`, schemaName+"_"+field.Type.Name())

			} else if arg.Value.Raw == "one2many" {

				gormType = fmt.Sprintf(`gorm:"foreignKey:ID;references:%s;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`,strings.Title(field.Name))

			}else if arg.Value.Raw == "one2one" {

				gormType = fmt.Sprintf(`gorm:"foriegnKey:%s;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`, strings.Title(field.Name))

			}

		} else {

			gormType = fmt.Sprintf(`gorm:"column:%s"`, field.Name)

		}

	}

	directive = field.Directives.ForName("constraint")
	if directive != nil {

		arg := directive.Arguments.ForName("type")

		if arg != nil {

			if arg.Value.Raw == "check" {

				value := directive.Arguments.ForName("value").Value.Raw

				gormType = fmt.Sprintf(`gorm:"check:%s"`, value)

			} else if arg.Value.Raw == "default" {

				value := directive.Arguments.ForName("value").Value.Raw

				gormType = fmt.Sprintf(`gorm:"default:%s"`, value)

			} else if arg.Value.Raw == "unique" {

				gormType = fmt.Sprintf(`gorm:"column:%s;uniqueIndex"`, field.Name)

			} else if arg.Value.Raw == "index" {

				gormType = `gorm:"index"`

			} else if arg.Value.Raw == "primarykey" {

				gormType = `gorm:"autoIncrement; primaryKey"`

			} else if arg.Value.Raw == "notnull" {

				gormType = `gorm:"not null"`

			}else if arg.Value.Raw == "autoincrement" {
				gormType = fmt.Sprintf(`gorm:"autoIncrement"`)

			}

		} else {

			gormType = fmt.Sprintf(`gorm:"column:%s"`, field.Name)

		}

	}

	return gormType

}