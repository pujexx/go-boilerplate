package gen

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)



func getSize(t string) *string{
	re := regexp.MustCompile(`(?m)\d+`)
	matches := re.FindAllString(t,1)
	if len(matches) > 0 {
		return &matches[0]
	}
	return nil
}

func generateCode(column Column,t string) string {
	validate := []string{}
	ifnull := ""
	if column.Null == "NO"{
		ifnull += "NOT NULL;"
		validate = append(validate,"required")
	}

	size := ""
	if s := getSize(column.Type);s != nil {
		size += "size:"+*s+";"
		if t == "string" {
			validate = append(validate,"max="+*s)
		}

	}

	if strings.Contains(strings.ToUpper(column.Field),"EMAIL") {
		validate = append(validate,"email")
	}

	if strings.Contains(t,"int") || strings.Contains(t,"double") {
		validate = append(validate,"number")
	}

	if strings.Contains(t,"bool") {
		validate = append(validate,"oneof=true false")
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
	matches := re.FindAllString(column.Type,1)
	if len(matches) > 0 {
		size += "type:"+matches[0]+";"
	}

	titles := strings.Split(column.Field,"_")
	title := strings.Title(strings.Join(titles," "))
	title = strings.Replace(title," ","",-1)
	validations := ""
	if len(validate) > 0 {
		validations = fmt.Sprintf(" validate:\"%v\"",strings.Join(validate,","))
	}
	code := title +" "+t+" `gorm:\"column:"+column.Field+";"+size+"\" json:\""+column.Field+"\""+validations+"`"
	return code
}

func ConverterField(column Column) string{
	Type := strings.ToUpper(column.Type)
	code := "\t"
	//String data types:
	if strings.HasPrefix(Type,"CHAR") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"VARCHAR") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"BINARY") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"VARBINARY") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"TINYBLOB") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"TINYTEXT") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"TEXT") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"BLOB") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"MEDIUMTEXT") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"MEDIUMTEXT") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"MEDIUMBLOB") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"LONGTEXT") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"LONGBLOB") {
		code += generateCode(column,"string")
	}
	if strings.HasPrefix(Type,"ENUM") {
		code += generateCode(column,"string")
	}

	if strings.HasPrefix(Type,"SET") {
		code += generateCode(column,"string")
	}

	//numeric data type
	if strings.HasPrefix(Type,"BIT") {
		code += generateCode(column,"int")
	}
	if strings.HasPrefix(Type,"TINYINT") {
		code += generateCode(column,"int")
	}
	if strings.HasPrefix(Type,"BOOL") {
		code += generateCode(column,"bool")
	}
	if strings.HasPrefix(Type,"BOOLEAN") {
		code += generateCode(column,"bool")
	}
	if strings.HasPrefix(Type,"SMALLINT") {
		code += generateCode(column,"bool")
	}
	if strings.HasPrefix(Type,"MEDIUMINT") {
		code += generateCode(column,"int")
	}
	if strings.HasPrefix(Type,"INT") {
		code += generateCode(column,"int")
	}
	if strings.HasPrefix(Type,"INTEGER") {
		code += generateCode(column,"int")
	}
	if strings.HasPrefix(Type,"BIGINT") {
		code += generateCode(column,"int64")
	}
	if strings.HasPrefix(Type,"FLOAT") {
		code += generateCode(column,"float")
	}
	if strings.HasPrefix(Type,"DOUBLE") {
		code += generateCode(column,"double")
	}
	if strings.HasPrefix(Type,"DECIMAL") {
		code += generateCode(column,"double")
	}

	if strings.HasPrefix(Type,"DEC") {
		code += generateCode(column,"double")
	}



	if strings.HasPrefix(Type,"DATE") {
		code += generateCode(column,"time.Time")
	}

	if strings.HasPrefix(Type,"TIMESTAMP") {
		code += generateCode(column,"time.Time")
	}
	if strings.HasPrefix(Type,"TIME") {
		code += generateCode(column,"time.Time")
	}
	if strings.HasPrefix(Type,"YEAR") {
		code += generateCode(column,"time.Time")
	}

	return code+"\n"

}

type domainGenerator struct {
	gen GeneratorField
}

func NewDomainGenerator(gen GeneratorField) *domainGenerator{
	return &domainGenerator{
		gen: gen,
	}
}

func (n domainGenerator)Generator(){

	result := n.gen.Tables()

	for _,r := range result {
		columns := n.gen.Columns(r)
		title := n.gen.TitleName(r)
		templateFile , err :=ioutil.ReadFile("lib/gen/assets/domain_template.txt")

		if err != nil {
			fmt.Println(err)
		}

		template := string(templateFile)

		tempBody := ""
		tempImport := []string{}
		for _, column := range columns {
			if strings.HasPrefix(strings.ToUpper(column.Type),"DATE") || strings.HasPrefix(strings.ToUpper(column.Type),"TIME")  {
				//tempImport = "\"time\"\n"
				tempImport = append(tempImport,"\t\"time\"")
			}
			tempBody += ConverterField(column)
		}

		imports := uniqueStrings(tempImport)
		imports = append(imports,fmt.Sprintf("\t\"%v/lib\"",n.gen.ModuleName()))
		template = strings.Replace(template,"{{import}}",strings.Join(imports,"\n"),-1)
		template = strings.Replace(template,"{{domain_struct}}",title,-1)
		template = strings.Replace(template,"{{fields}}",tempBody,-1)
		template = strings.Replace(template,"{{table_name}}",r,-1)

		fmt.Println(string(template))
		ioutil.WriteFile("domain/"+r+".go",[]byte(template),0644)
		fmt.Println("==========================")
	}
}

func uniqueStrings(stringslice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringslice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}