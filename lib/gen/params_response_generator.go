package gen

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func getSizeParameter(t string) *string {
	re := regexp.MustCompile(`(?m)\d+`)
	matches := re.FindAllString(t, 1)
	if len(matches) > 0 {
		return &matches[0]
	}
	return nil
}

func generateCodeParameter(column Column, t string) string {
	validate := []string{}
	ifnull := ""
	if column.Null == "NO" {
		ifnull += "NOT NULL;"
		validate = append(validate, "required")
	}

	size := ""
	if s := getSizeParameter(column.Type); s != nil {
		size += "size:" + *s + ";"
		if t == "string" {
			validate = append(validate, "max="+*s)
		}

	}

	if strings.Contains(strings.ToUpper(column.Field), "EMAIL") {
		validate = append(validate, "email")
	}

	if strings.Contains(t, "int") || strings.Contains(t, "double") {
		validate = append(validate, "number")
	}

	if strings.Contains(t, "bool") {
		validate = append(validate, "oneof=true false")
	}

	if strings.Contains(t, "time.Time") {
		validate = append(validate, "datetime=2006-01-02")
	}

	if column.Key == "PRI" {
		size += "primaryKey;"
		validate = []string{}
	}

	if column.Extra == "auto_increment" {
		size += "autoIncrement;"
		validate = []string{}
	}

	var re = regexp.MustCompile(`\b[^\d\W]+\b`)
	matches := re.FindAllString(column.Type, 1)
	if len(matches) > 0 {
		size += "type:" + matches[0] + ";"
	}

	titles := strings.Split(column.Field, "_")
	title := strings.Title(strings.Join(titles, " "))
	title = strings.Replace(title, " ", "", -1)
	validations := ""
	if len(validate) > 0 {
		validations = fmt.Sprintf(" validate:\"%v\"", strings.Join(validate, ","))
	}
	if t == "time.Time" {
		t = "string"
	}
	code := title + " " + t + " `json:\"" + column.Field + "\"" + validations + "`"
	return code
}

func ConverterFieldParameter(column Column) string {
	Type := strings.ToUpper(column.Type)
	code := "\t"
	//String data types:
	if strings.HasPrefix(Type, "CHAR") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "VARCHAR") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "BINARY") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "VARBINARY") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "TINYBLOB") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "TINYTEXT") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "TEXT") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "BLOB") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "MEDIUMTEXT") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "MEDIUMTEXT") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "MEDIUMBLOB") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "LONGTEXT") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "LONGBLOB") {
		code += generateCodeParameter(column, "string")
	}
	if strings.HasPrefix(Type, "ENUM") {
		code += generateCodeParameter(column, "string")
	}

	if strings.HasPrefix(Type, "SET") {
		code += generateCodeParameter(column, "string")
	}

	//numeric data type
	if strings.HasPrefix(Type, "BIT") {
		code += generateCodeParameter(column, "int")
	}
	if strings.HasPrefix(Type, "TINYINT") {
		code += generateCodeParameter(column, "int")
	}
	if strings.HasPrefix(Type, "BOOL") {
		code += generateCodeParameter(column, "bool")
	}
	if strings.HasPrefix(Type, "BOOLEAN") {
		code += generateCodeParameter(column, "bool")
	}
	if strings.HasPrefix(Type, "SMALLINT") {
		code += generateCodeParameter(column, "bool")
	}
	if strings.HasPrefix(Type, "MEDIUMINT") {
		code += generateCodeParameter(column, "int")
	}
	if strings.HasPrefix(Type, "INT") {
		code += generateCodeParameter(column, "int")
	}
	if strings.HasPrefix(Type, "INTEGER") {
		code += generateCodeParameter(column, "int")
	}
	if strings.HasPrefix(Type, "BIGINT") {
		code += generateCodeParameter(column, "int64")
	}
	if strings.HasPrefix(Type, "FLOAT") {
		code += generateCodeParameter(column, "float")
	}
	if strings.HasPrefix(Type, "DOUBLE") {
		code += generateCodeParameter(column, "double")
	}
	if strings.HasPrefix(Type, "DECIMAL") {
		code += generateCodeParameter(column, "double")
	}

	if strings.HasPrefix(Type, "DEC") {
		code += generateCodeParameter(column, "double")
	}

	if strings.HasPrefix(Type, "DATE") {
		code += generateCodeParameter(column, "time.Time")
	}

	if strings.HasPrefix(Type, "TIMESTAMP") {
		code += generateCodeParameter(column, "time.Time")
	}
	if strings.HasPrefix(Type, "TIME") {
		code += generateCodeParameter(column, "time.Time")
	}
	if strings.HasPrefix(Type, "YEAR") {
		code += generateCodeParameter(column, "time.Time")
	}

	return code + "\n"

}

type paramGenerator struct {
	gen GeneratorField
}

func NewParameterGenerator(gen GeneratorField) *paramGenerator {
	return &paramGenerator{
		gen: gen,
	}
}

func (n paramGenerator) GeneratorParameter() {

	result := n.gen.Tables()

	for _, r := range result {
		columns := n.gen.Columns(r)
		title := n.gen.TitleName(r)
		templateFile, err := ioutil.ReadFile("lib/gen/assets/parameter_response_template.txt")

		if err != nil {
			fmt.Println(err)
		}

		template := string(templateFile)

		tempBody := ""
		tempImport := []string{}
		for _, column := range columns {
			if strings.HasPrefix(strings.ToUpper(column.Type), "DATE") || strings.HasPrefix(strings.ToUpper(column.Type), "TIME") {
				//tempImport = "\"time\"\n"
				tempImport = append(tempImport, "\t\"time\"")
			}
			tempBody += ConverterFieldParameter(column)
		}

		imports := uniqueStrings(tempImport)
		imports = append(imports, fmt.Sprintf("\t\"%v/lib\"", n.gen.ModuleName()))
		template = strings.Replace(template, "{{import}}", strings.Join(imports, "\n"), -1)
		template = strings.Replace(template, "{{domain_struct}}", title, -1)
		template = strings.Replace(template, "{{fields}}", tempBody, -1)
		template = strings.Replace(template, "{{table_name}}", r, -1)

		fmt.Println(string(template))
		ioutil.WriteFile(r+"/handler/http/"+r+"_models.go", []byte(template), 0644)
		fmt.Println("==========================")
	}
}
