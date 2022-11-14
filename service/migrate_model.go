package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (s *service) MigrateModels() {
	file, err := os.Open("graph/model/models_gen.go")
	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	// `create table if not exists Home (
	// 	id serial not null primary key,
	// 	name varchar(255),
	// 	value double precision,
	// 	data date
	//  );`

	// alter table tweet
	//   add column "user" text,
	//   add constraint fk_test
	//   foreign key ("user")
	//   references "User" (id);

	m := make(map[string][]string)
	relation_map := make(map[string][]string)
	query := "create table if not exists "
	sql_table_name := ""
	sql_column_name := ""
	for scanner.Scan() {
		temp := scanner.Text()

		// Table Name
		if strings.Contains(temp, "type") && strings.Contains(temp, "struct") {
			table_name := strings.Split(temp, " ")
			text = append(text, table_name[1])

			// check for Table name user(Reserved keyword)
			// if (table_name[1] == "User") || (table_name[1] == "USER") || (table_name[1] == "user") {
			// 	query = query + "\"" + table_name[1] + "\"" + "( "
			// } else {
			query = query + "public." + table_name[1] + "( "
			// }
			sql_table_name = "public." + table_name[1]

		} else if strings.Contains(temp, "json") {

			// Column Name
			temp1 := strings.Split(temp, "`json:")
			x := temp1[1]

			if (strings.Contains(temp, "int")) || (strings.Contains(temp, "string")) || (strings.Contains(temp, "[]") || (strings.Contains(temp, "bool"))) {
				query = query + x[1:len(x)-2]
				sql_column_name = x[1 : len(x)-2]
			} else {
				sql_column_name = x[1 : len(x)-2]
			}

			// Data Type
			y := temp1[0]
			if strings.Contains(y, "int") {
				if !(strings.Contains(query, "NUMERIC") || strings.Contains(query, "TEXT") || strings.Contains(query, "JSON") || strings.Contains(query, "BOOLEAN")) {
					query = query + " NUMERIC PRIMARY KEY, "
				} else {
					query = query + " NUMERIC, "
				}
			} else if strings.Contains(y, "string") {
				if !(strings.Contains(query, "NUMERIC") || strings.Contains(query, "TEXT") || strings.Contains(query, "JSON") || strings.Contains(query, "BOOLEAN")) {
					query = query + " TEXT PRIMARY KEY, "
				} else {
					query = query + " TEXT, "
				}
			} else if strings.Contains(y, "[]*") {
				query = query + " JSON[], "
			} else if strings.Contains(y, "bool") {
				query = query + " BOOLEAN, "
			} else if strings.Contains(y, "*") && (!(strings.Contains(y, "int")) || (strings.Contains(y, "string")) || (strings.Contains(y, "[]") || (strings.Contains(y, "bool")))) {
				// query = query + " JSON, "
				// fmt.Println(y)
				temp1Arr := strings.Split(y, "*")
				mappedField := temp1Arr[1]
				mappedField = strings.ReplaceAll(mappedField, " ", "")
				// Defining map with array of fields
				_, ok := m[sql_table_name]
				_, okforRel := relation_map[sql_table_name]
				if ok && okforRel {
					list_of_fields := m[sql_table_name]
					list_of_fields = append(list_of_fields, sql_column_name)
					m[sql_table_name] = list_of_fields

					list_of_mappedWith := relation_map[sql_table_name]
					list_of_mappedWith = append(list_of_mappedWith, "public."+mappedField)
					relation_map[sql_table_name] = list_of_mappedWith

				} else {
					var list_of_fields []string
					list_of_fields = append(list_of_fields, sql_column_name)
					m[sql_table_name] = list_of_fields

					var list_of_mappedWith []string
					list_of_mappedWith = append(list_of_mappedWith, "public."+mappedField)
					relation_map[sql_table_name] = list_of_mappedWith
				}

			} else {
			}

			// Reseting the query and len variable for tracking Primary Key
		} else if strings.Contains(temp, "}") {
			query = query[:len(query)-2] + " );"
			if strings.Contains(query, "interface") {
				fmt.Println("Interface Not Allowed..")
			} else {
				// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				// defer cancelfunc()
				fmt.Println(query)
				err := s.Repo.ExecQuery(query)
				// fmt.Println(res)
				if err != nil {
					log.Printf("Error %s when creating product table", err)
				}
			}
			query = "create table if not exists "
			sql_table_name = ""
			sql_column_name = ""
		} else {
		}
	}
	fmt.Println("Migrated Successfully..")

	//ALter Table
	/*for index, element := range m {
		alterQuery := "alter table public."
		x := relation_map[index]
		// alterQuery = alterQuery + checkReserveKeyword(index)
		alterQuery = alterQuery + index
		for id, record := range element {
			alterQuery = alterQuery + " add column "
			// target_element := checkReserveKeyword(record)
			// target_relation_field := checkReserveKeyword(x[id])

			target_element := record
			target_relation_field := x[id]
			alterQuery = alterQuery + target_element + " text, add constraint fk" + target_element + " foreign key (" + target_element + ") references " + target_relation_field + " (id) ,"
		}
		alterQuery = alterQuery[:len(alterQuery)-1]
		alterQuery = alterQuery + ";"
		// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancelfunc()
		fmt.Println(alterQuery)
		err := s.Repo.ExecQuery(alterQuery)
		// fmt.Println(res)
		if err != nil {
			log.Printf("Error %s when creating product table", err)
		}
		alterQuery = "alter table public."

	}*/

}

/*func checkReserveKeyword(s string) string {
	if s == "user" {
		return "\"user\""
	} else if s == "User" {
		return "\"User\""
	} else if s == "USER" {
		return "\"USER\""
	}
	return s
}*/

// func (service *service) MigrateModels() {

// 	fileName := "./graph/model/models_gen.go"
// 	modelFile, err := os.Open(fileName)
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer modelFile.Close()

// 	byteValue, _ := ioutil.ReadAll(modelFile)

// 	fset := token.NewFileSet()

// 	f, err := parser.ParseFile(fset, fileName, byteValue, parser.ParseComments)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// ast.Print(fset, f)

// 	// return

// 	// for k, _ := range f.Scope.Objects {
// 	// 	log.Println(f.Scope.Outer.Lookup(k))
// 	// 	// service.Repo.MigrateSchema(k, v)
// 	// }

// 	// return

// 	structMap := make(map[string]interface{}, 0)
// 	associationMap := make(map[string]string, 0)
// 	ast.Inspect(f, func(n ast.Node) bool {

// 		switch t := n.(type) {
// 		case *ast.TypeSpec:

// 			log.Println("dec-->", t.Name.Obj)
// 			structType := t.Type.(*ast.StructType)
// 			structFields := []reflect.StructField{}

// 			for _, field := range structType.Fields.List {

// 				star, ok := field.Type.(*ast.StarExpr)
// 				if ok {
// 					i, _ := star.X.(*ast.Ident)
// 					if val, present := structMap[strings.ToLower(i.Name)]; present {
// 						for _, name := range field.Names {
// 							fmt.Printf("\tField: name=%s type=%s\n", name.Name, val)

// 							associationMap[strings.ToLower(t.Name.Name)] = strings.ToLower(name.Name)

// 							// structFields = append(structFields, reflect.StructField{

// 							// 	Name: name.Name,
// 							// 	Type: reflect.ValueOf(val).Elem().Addr().Type(),
// 							// 	Tag:  reflect.StructTag(fmt.Sprintf(`gorm:"foreignKey:%s"`, name.Name+"ID")),
// 							// })

// 							// structFields = append(structFields, reflect.StructField{

// 							// 	Name: name.Name + "ID",
// 							// 	Type: reflect.TypeOf(""),
// 							// 	Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, name.Name+"ID")),
// 							// })
// 						}

// 					} else {

// 					}

// 				}

// 				i, ok := field.Type.(*ast.Ident)
// 				if ok {
// 					fieldType := i.Name
// 					for _, name := range field.Names {
// 						fmt.Printf("\tField: name=%s type=%s\n", name.Name, fieldType)
// 						structFields = append(structFields, reflect.StructField{

// 							Name: name.Name,
// 							Type: getType(fieldType),
// 							Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, name.Name)),
// 						})
// 					}
// 				}

// 			}

// 			// structFields = append(structFields, reflect.StructField{Name: strings.ToLower(t.Name.Name),
// 			// 	Type:      reflect.TypeOf(""),
// 			// 	Tag:       reflect.StructTag(`json:"@type" gorm:"-"`),
// 			// 	Anonymous: true,
// 			// })

// 			ss := reflect.StructOf(structFields)
// 			d := reflect.New(ss).Elem().Interface()

// 			structMap[strings.ToLower(t.Name.Name)] = d
// 			log.Println(t.Name.Name, "-->", d)
// 			service.Repo.MigrateSchema(strings.ToLower(t.Name.Name), d)
// 			// if err == nil {
// 			// 	structMap[strings.ToLower(t.Name.Name)] = vaue
// 			// }
// 		}

// 		return true
// 	})

// 	for k, v := range associationMap {
// 		model := structMap[k]
// 		service.Repo.AlterTable(model, v)
// 	}

// 	log.Println(structMap)

// }

// func getType(datatype string) reflect.Type {

// 	// log.Println(datatype)
// 	switch datatype {
// 	case "string":
// 		return reflect.TypeOf("string")
// 	case "int":
// 		return reflect.TypeOf(1)
// 	case "float64":
// 		return reflect.TypeOf(float64(0))
// 	case "date":
// 		return reflect.TypeOf(time.Time{})
// 	case "time":
// 		return reflect.TypeOf(time.Time{})

// 	}

// 	return reflect.TypeOf("string")
// }