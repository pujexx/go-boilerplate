package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pujexx/go-boilerplate/lib"
	"github.com/pujexx/go-boilerplate/lib/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"

)





func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	user := flag.String("user",os.Getenv("DB_USER"),"user")
	password := flag.String("password",os.Getenv("DB_PASSWORD"),"password")
	host := flag.String("host",os.Getenv("DB_HOST"),"Host")
	dbname := flag.String("db",os.Getenv("DB_NAME"),"DB name")
	fmt.Println(*user)
	flag.Parse()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	con := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",*user,*password, *host,"3306",*dbname)
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{Logger : newLogger })

	if err != nil {
		panic("error connection database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)


	r := mux.NewRouter()
	pathPrefix := r.PathPrefix("/v1").Subrouter()

	pathPrefix.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		lib.BaseResponse(lib.Response{
			Code: "200",
			Data: "pingpong",
		},writer,request)
		return
	})

	//implement generator

	//end of implement generator


	generator := gen.NewGenerator(db)
	generator.GeneratorCLI()
	http.ListenAndServe(":8000",r)
	fmt.Println("service on")

}
