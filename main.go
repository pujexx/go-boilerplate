package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pujexx/go-boilerplate/infrastructure"
	"github.com/pujexx/go-boilerplate/lib/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"

)




// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1
// @query.collection.format multi

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

	user := flag.String("user", os.Getenv("DB_USER"), "user")
	password := flag.String("password", os.Getenv("DB_PASSWORD"), "password")
	host := flag.String("host", os.Getenv("DB_HOST"), "Host")
	dbname := flag.String("db", os.Getenv("DB_NAME"), "DB name")
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
	con := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", *user, *password, *host, "3306", *dbname)
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{Logger: newLogger})

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

	infrastructure.NewPingHandler(pathPrefix)

	//implement generator

	//end of implement generator

	generator := gen.NewGenerator(db)
	generator.GeneratorCLI()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			next.ServeHTTP(w, r)
		})
	})
	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
	fmt.Println("service on")

}
