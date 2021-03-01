package gen

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Column struct {
	Field string `gorm:"column:Field"`
	Type string `gorm:"column:Type"`
	Null string `gorm:"column:Null"`
	Key string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra string `gorm:"column:Extra"`

}

type GeneratorField struct {
	DB *gorm.DB
}

func NewGenerator(db * gorm.DB)*GeneratorField{
	return &GeneratorField{
		DB: db,
	}
}


func (g GeneratorField)Tables() []string{
	var result []string
	g.DB.Debug().Raw("SHOW TABLES").Scan(&result)
	return result
}

func (g GeneratorField)Columns(table string) []Column{
	var columns []Column
	g.DB.Raw(fmt.Sprintf("SHOW COLUMNS FROM %v",table)).Scan(&columns)
	return columns
}

func (g GeneratorField) PrimaryKey(table string) string{
	var column Column
	g.DB.Raw(fmt.Sprintf("SHOW COLUMNS FROM %v",table)).Scan(&column)

	return column.Field
}

func (g GeneratorField) PrimaryKeyTitle(table string) string{
	titles := strings.Split(g.PrimaryKey(table),"_")
	title := strings.Title(strings.Join(titles," "))
	title = strings.Replace(title," ","",-1)
	return title
}

func (g GeneratorField) ModuleName()string{
	gomodfile , err :=ioutil.ReadFile("go.mod")
	if err != nil {
		fmt.Println(err)
	}
	gomod := string(gomodfile)
	var re = regexp.MustCompile(`(?m)^(module\s)(.*)`)
	match := re.FindStringSubmatch(gomod)
	if len(match) > 2 {
		return match[2]
	}
	return ""
}

func (g GeneratorField) TitleName(table string) string{
	titles := strings.Split(table,"_")
	title := strings.Title(strings.Join(titles," "))
	title = strings.Replace(title," ","",-1)
	return title
}

func (g GeneratorField) TableName(table string) string{
	titles := strings.Split(table,"_")
	title := strings.Join(titles," ")
	title = strings.Replace(title," ","",-1)
	return title
}

func (g GeneratorField) InitalTableName(table string) string{
	titles := []string{}
	for _, t := range strings.Split(table,"_") {
		titles = append(titles,strings.ToLower( string(t[0])) )
	}
	title := strings.Join(titles,"")
	return title
}

func (g GeneratorField)Generator(){
	g.GenerateCRUD()
}

func (g GeneratorField)GeneratorCLI(){
	app := cli.NewApp()
	app.EnableBashCompletion = true


	app.Commands = []*cli.Command{
		{
			Name:                   "grill",
			Aliases:                []string{"g"},
			Usage:                  "the generator code",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "generate all code",
					Action: func(c *cli.Context) error {
						g.Generator()
						os.Exit(0)
						return nil
					},
				},
				{
					Name:  "domain",
					Usage: "generate domains code",
					Action: func(c *cli.Context) error {
						g.GeneratorDomain()
						os.Exit(0)
						return nil
					},
				},
				{
					Name:  "repository",
					Usage: "generate repository code",
					Action: func(c *cli.Context) error {
						g.GeneratorRepository()
						os.Exit(0)
						return nil
					},
				},
				{
					Name:  "service",
					Usage: "generate service code",
					Action: func(c *cli.Context) error {
						g.GeneratorService()
						os.Exit(0)
						return nil
					},
				},
				{
					Name:  "http-handler",
					Usage: "generate service code",
					Action: func(c *cli.Context) error {
						g.GeneratorHttpHandler()
						os.Exit(0)
						return nil
					},
				},
				{
					Name:  "implement",
					Usage: "generate implement code",
					Action: func(c *cli.Context) error {
						g.GeneratorImplementMain()
						os.Exit(0)
						return nil
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (g GeneratorField) GenerateCRUD() {
	g.GeneratorDomain()
	g.GeneratorRepository()
	g.GeneratorService()
	g.GeneratorHttpHandler()
	g.GeneratorImplementMain()
}


func (g GeneratorField) GeneratorDomain()  {
	if _, err := os.Stat("domain"); os.IsNotExist(err) {
		fmt.Println("create directory ", "domain")
		os.Mkdir("domain",0755)
	}
	dom := NewDomainGenerator(g)
	dom.Generator()

}


func (g GeneratorField) GeneratorRepository()  {
	g.CreateDirectory()
	for _,table := range g.Tables() {
		g.GenerateRepository(table)
	}

}

func (g GeneratorField) GeneratorService()  {
	g.CreateDirectory()
	for _,table := range g.Tables() {
		g.GenerateService(table)
	}

}

func (g GeneratorField) GeneratorHttpHandler()  {
	g.CreateDirectory()
	for _,table := range g.Tables() {
		g.GenerateHttpHandler(table)
	}

}

func (g GeneratorField) GeneratorImplementMain()  {
	g.ImplementMain()
}

func (g GeneratorField) CreateDirectory(){
	for _,table := range g.Tables() {
		if _, err := os.Stat(table); os.IsNotExist(err) {
			fmt.Println("create directory ", table)
			os.Mkdir(table,0755)
		}
		if _, err := os.Stat(fmt.Sprintf("%v/repository",table)); os.IsNotExist(err) {
			fmt.Println("create directory repository", table)
			os.Mkdir(fmt.Sprintf("%v/repository",table),0755)
		}
		if _, err := os.Stat(fmt.Sprintf("%v/service",table)); os.IsNotExist(err) {
			fmt.Println("create directory service", table)
			os.Mkdir(fmt.Sprintf("%v/service",table),0755)
		}

		if _, err := os.Stat(fmt.Sprintf("%v/handler",table)); os.IsNotExist(err) {
			fmt.Println("create directory handler", table)
			os.Mkdir(fmt.Sprintf("%v/handler",table),0755)
		}

		if _, err := os.Stat(fmt.Sprintf("%v/handler/http",table)); os.IsNotExist(err) {
			fmt.Println("create directory handler http", table)
			os.Mkdir(fmt.Sprintf("%v/handler/http",table),0755)
		}
	}
}


func (g GeneratorField) GenerateRepository(table string) {
	templateFile , err :=ioutil.ReadFile("lib/gen/assets/repository_template.txt")
	if err != nil {
		fmt.Println(err)
	}
	template := string(templateFile)
	template = strings.Replace(template,"{{table_name}}", table,-1)
	template = strings.Replace(template,"{{module_name}}", g.ModuleName(),-1)
	template = strings.Replace(template,"{{import}}","",-1)
	template = strings.Replace(template,"{{table_title}}",g.TableName(table),-1)
	template = strings.Replace(template,"{{domain_struct}}",g.TitleName(table),-1)
	template = strings.Replace(template,"{{initial_table}}",g.InitalTableName(table),-1)
	template = strings.Replace(template,"{{primary_key}}",g.PrimaryKey(table),-1)


	fileName := fmt.Sprintf("%v/repository/%v_repository.go",table,table)
	fmt.Println(fileName)

	if err := ioutil.WriteFile(fileName,[]byte(template),0644);err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========================")

}


func (g GeneratorField) GenerateService(table string) {
	templateFile , err :=ioutil.ReadFile("lib/gen/assets/service_template.txt")
	if err != nil {
		fmt.Println(err)
	}
	template := string(templateFile)
	template = strings.Replace(template,"{{table_name}}", table,-1)
	template = strings.Replace(template,"{{module_name}}", g.ModuleName(),-1)
	template = strings.Replace(template,"{{import}}","",-1)
	template = strings.Replace(template,"{{table_title}}",g.TableName(table),-1)
	template = strings.Replace(template,"{{domain_struct}}",g.TitleName(table),-1)
	template = strings.Replace(template,"{{initial_table}}",g.InitalTableName(table),-1)
	template = strings.Replace(template,"{{primary_key}}",g.PrimaryKey(table),-1)


	fileName := fmt.Sprintf("%v/service/%v_service.go",table,table)
	fmt.Println(fileName)

	if err := ioutil.WriteFile(fileName,[]byte(template),0644);err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========================")

}



func (g GeneratorField) GenerateHttpHandler(table string) {
	templateFile , err :=ioutil.ReadFile("lib/gen/assets/http_handler_template.txt")
	if err != nil {
		fmt.Println(err)
	}
	template := string(templateFile)
	template = strings.Replace(template,"{{table_name}}", table,-1)
	template = strings.Replace(template,"{{module_name}}", g.ModuleName(),-1)
	template = strings.Replace(template,"{{import}}","",-1)
	template = strings.Replace(template,"{{table_title}}",g.TableName(table),-1)
	template = strings.Replace(template,"{{domain_struct}}",g.TitleName(table),-1)
	template = strings.Replace(template,"{{initial_table}}",g.InitalTableName(table),-1)
	template = strings.Replace(template,"{{primary_key}}",g.PrimaryKey(table),-1)
	template = strings.Replace(template,"{{primary_key_title}}",g.PrimaryKeyTitle(table),-1)


	fileName := fmt.Sprintf("%v/handler/http/%v_http.go",table,table)
	fmt.Println(fileName)

	if err := ioutil.WriteFile(fileName,[]byte(template),0644);err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========================")

}

func (g GeneratorField) ImplementMain() {
	templateFile , err :=ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println(err)
	}
	template := string(templateFile)
	imp := getImport(template)
	if imp == nil {
		fmt.Println("import not found")
	}
	impr := *imp
	template = strings.Replace(template,*imp, "{{import}}" ,-1)
	impr = strings.Replace(impr,")","",-1)

	importVar := []string{}
	implementVar := []string{}
	implementVar = append(implementVar,"//implement generator")
	for _,table := range g.Tables() {
		importVar = append(importVar,fmt.Sprintf("\t_%vRepo \"%v/%v/repository\"",table,g.ModuleName(),table))
		importVar = append(importVar,fmt.Sprintf("\t_%vService \"%v/%v/service\"",table,g.ModuleName(),table))
		importVar = append(importVar,fmt.Sprintf("\t_%vHttp \"%v/%v/handler/http\"",table,g.ModuleName(),table))

		implementVar = append(implementVar,fmt.Sprintf("\t//=====implement %v =====",table))
		implementVar = append(implementVar,fmt.Sprintf("\t%vRepo := _%vRepo.New%vRepository(db)",table,table,g.TitleName(table)))
		implementVar = append(implementVar,fmt.Sprintf("\t%vService := _%vService.New%vService(%vRepo)",table,table,g.TitleName(table),table))
		implementVar = append(implementVar,fmt.Sprintf("\t_%vHttp.New%vHttpHandler(pathPrefix,%vService)",table,g.TitleName(table),table))

	}
	implementVar = append(implementVar,"\t//end of implement generator")
	impr += "\n"+strings.Join(importVar,"\n")+"\n)\n"
	template = strings.Replace(template,"{{import}}",impr,-1)

	imple := getImplement(template)
	if imple == nil {
		fmt.Println("implement not found")
	}
	template = strings.Replace(template,*imple, "{{implement}}" ,-1)

	template = strings.Replace(template,"{{implement}}", strings.Join(implementVar,"\n"),-1)

	fileName := "main.go"
	fmt.Println(fileName)

	if err := ioutil.WriteFile(fileName,[]byte(template),0644);err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========================")

}


func getImport(t string) *string{
	re := regexp.MustCompile(`(?m)import\s*(\((?:\[??[^\[]*?)\))`)
	matches := re.FindAllString(t,1)
	if len(matches) > 0 {
		return &matches[0]
	}
	return nil
}

func getImplement(t string) *string{
	re := regexp.MustCompile(`(?m)(\/\/implement generator(?:\[??[^\[]*?)\/\/end of implement generator)`)
	matches := re.FindAllString(t,1)
	if len(matches) > 0 {
		return &matches[0]
	}
	return nil
}

