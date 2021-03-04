// Package classification awesome.
//
// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Supun Muthutantri<mydocs@example.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta

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

	_access_rolesHttp "github.com/pujexx/go-boilerplate/access_roles/handler/http"
	_access_rolesRepo "github.com/pujexx/go-boilerplate/access_roles/repository"
	_access_rolesService "github.com/pujexx/go-boilerplate/access_roles/service"
	_categoriesHttp "github.com/pujexx/go-boilerplate/categories/handler/http"
	_categoriesRepo "github.com/pujexx/go-boilerplate/categories/repository"
	_categoriesService "github.com/pujexx/go-boilerplate/categories/service"
	_cdnsHttp "github.com/pujexx/go-boilerplate/cdns/handler/http"
	_cdnsRepo "github.com/pujexx/go-boilerplate/cdns/repository"
	_cdnsService "github.com/pujexx/go-boilerplate/cdns/service"
	_extra_ordersHttp "github.com/pujexx/go-boilerplate/extra_orders/handler/http"
	_extra_ordersRepo "github.com/pujexx/go-boilerplate/extra_orders/repository"
	_extra_ordersService "github.com/pujexx/go-boilerplate/extra_orders/service"
	_featuresHttp "github.com/pujexx/go-boilerplate/features/handler/http"
	_featuresRepo "github.com/pujexx/go-boilerplate/features/repository"
	_featuresService "github.com/pujexx/go-boilerplate/features/service"
	_inventoriesHttp "github.com/pujexx/go-boilerplate/inventories/handler/http"
	_inventoriesRepo "github.com/pujexx/go-boilerplate/inventories/repository"
	_inventoriesService "github.com/pujexx/go-boilerplate/inventories/service"
	_ledgersHttp "github.com/pujexx/go-boilerplate/ledgers/handler/http"
	_ledgersRepo "github.com/pujexx/go-boilerplate/ledgers/repository"
	_ledgersService "github.com/pujexx/go-boilerplate/ledgers/service"
	_membersHttp "github.com/pujexx/go-boilerplate/members/handler/http"
	_membersRepo "github.com/pujexx/go-boilerplate/members/repository"
	_membersService "github.com/pujexx/go-boilerplate/members/service"
	_ordersHttp "github.com/pujexx/go-boilerplate/orders/handler/http"
	_ordersRepo "github.com/pujexx/go-boilerplate/orders/repository"
	_ordersService "github.com/pujexx/go-boilerplate/orders/service"
	_productsHttp "github.com/pujexx/go-boilerplate/products/handler/http"
	_productsRepo "github.com/pujexx/go-boilerplate/products/repository"
	_productsService "github.com/pujexx/go-boilerplate/products/service"
	_rolesHttp "github.com/pujexx/go-boilerplate/roles/handler/http"
	_rolesRepo "github.com/pujexx/go-boilerplate/roles/repository"
	_rolesService "github.com/pujexx/go-boilerplate/roles/service"
	_storesHttp "github.com/pujexx/go-boilerplate/stores/handler/http"
	_storesRepo "github.com/pujexx/go-boilerplate/stores/repository"
	_storesService "github.com/pujexx/go-boilerplate/stores/service"
	_supliersHttp "github.com/pujexx/go-boilerplate/supliers/handler/http"
	_supliersRepo "github.com/pujexx/go-boilerplate/supliers/repository"
	_supliersService "github.com/pujexx/go-boilerplate/supliers/service"
	_transactionsHttp "github.com/pujexx/go-boilerplate/transactions/handler/http"
	_transactionsRepo "github.com/pujexx/go-boilerplate/transactions/repository"
	_transactionsService "github.com/pujexx/go-boilerplate/transactions/service"
	_user_rolesHttp "github.com/pujexx/go-boilerplate/user_roles/handler/http"
	_user_rolesRepo "github.com/pujexx/go-boilerplate/user_roles/repository"
	_user_rolesService "github.com/pujexx/go-boilerplate/user_roles/service"
	_usersHttp "github.com/pujexx/go-boilerplate/users/handler/http"
	_usersRepo "github.com/pujexx/go-boilerplate/users/repository"
	_usersService "github.com/pujexx/go-boilerplate/users/service"
	_variant_optionsHttp "github.com/pujexx/go-boilerplate/variant_options/handler/http"
	_variant_optionsRepo "github.com/pujexx/go-boilerplate/variant_options/repository"
	_variant_optionsService "github.com/pujexx/go-boilerplate/variant_options/service"
	_variant_valuesHttp "github.com/pujexx/go-boilerplate/variant_values/handler/http"
	_variant_valuesRepo "github.com/pujexx/go-boilerplate/variant_values/repository"
	_variant_valuesService "github.com/pujexx/go-boilerplate/variant_values/service"
	_variantsHttp "github.com/pujexx/go-boilerplate/variants/handler/http"
	_variantsRepo "github.com/pujexx/go-boilerplate/variants/repository"
	_variantsService "github.com/pujexx/go-boilerplate/variants/service"
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
	//=====implement access_roles =====
	access_rolesRepo := _access_rolesRepo.NewAccessRolesRepository(db)
	access_rolesService := _access_rolesService.NewAccessRolesService(access_rolesRepo)
	_access_rolesHttp.NewAccessRolesHttpHandler(pathPrefix, access_rolesService)
	//=====implement categories =====
	categoriesRepo := _categoriesRepo.NewCategoriesRepository(db)
	categoriesService := _categoriesService.NewCategoriesService(categoriesRepo)
	_categoriesHttp.NewCategoriesHttpHandler(pathPrefix, categoriesService)
	//=====implement cdns =====
	cdnsRepo := _cdnsRepo.NewCdnsRepository(db)
	cdnsService := _cdnsService.NewCdnsService(cdnsRepo)
	_cdnsHttp.NewCdnsHttpHandler(pathPrefix, cdnsService)
	//=====implement extra_orders =====
	extra_ordersRepo := _extra_ordersRepo.NewExtraOrdersRepository(db)
	extra_ordersService := _extra_ordersService.NewExtraOrdersService(extra_ordersRepo)
	_extra_ordersHttp.NewExtraOrdersHttpHandler(pathPrefix, extra_ordersService)
	//=====implement features =====
	featuresRepo := _featuresRepo.NewFeaturesRepository(db)
	featuresService := _featuresService.NewFeaturesService(featuresRepo)
	_featuresHttp.NewFeaturesHttpHandler(pathPrefix, featuresService)
	//=====implement inventories =====
	inventoriesRepo := _inventoriesRepo.NewInventoriesRepository(db)
	inventoriesService := _inventoriesService.NewInventoriesService(inventoriesRepo)
	_inventoriesHttp.NewInventoriesHttpHandler(pathPrefix, inventoriesService)
	//=====implement ledgers =====
	ledgersRepo := _ledgersRepo.NewLedgersRepository(db)
	ledgersService := _ledgersService.NewLedgersService(ledgersRepo)
	_ledgersHttp.NewLedgersHttpHandler(pathPrefix, ledgersService)
	//=====implement members =====
	membersRepo := _membersRepo.NewMembersRepository(db)
	membersService := _membersService.NewMembersService(membersRepo)
	_membersHttp.NewMembersHttpHandler(pathPrefix, membersService)
	//=====implement orders =====
	ordersRepo := _ordersRepo.NewOrdersRepository(db)
	ordersService := _ordersService.NewOrdersService(ordersRepo)
	_ordersHttp.NewOrdersHttpHandler(pathPrefix, ordersService)
	//=====implement products =====
	productsRepo := _productsRepo.NewProductsRepository(db)
	productsService := _productsService.NewProductsService(productsRepo)
	_productsHttp.NewProductsHttpHandler(pathPrefix, productsService)
	//=====implement roles =====
	rolesRepo := _rolesRepo.NewRolesRepository(db)
	rolesService := _rolesService.NewRolesService(rolesRepo)
	_rolesHttp.NewRolesHttpHandler(pathPrefix, rolesService)
	//=====implement stores =====
	storesRepo := _storesRepo.NewStoresRepository(db)
	storesService := _storesService.NewStoresService(storesRepo)
	_storesHttp.NewStoresHttpHandler(pathPrefix, storesService)
	//=====implement supliers =====
	supliersRepo := _supliersRepo.NewSupliersRepository(db)
	supliersService := _supliersService.NewSupliersService(supliersRepo)
	_supliersHttp.NewSupliersHttpHandler(pathPrefix, supliersService)
	//=====implement transactions =====
	transactionsRepo := _transactionsRepo.NewTransactionsRepository(db)
	transactionsService := _transactionsService.NewTransactionsService(transactionsRepo)
	_transactionsHttp.NewTransactionsHttpHandler(pathPrefix, transactionsService)
	//=====implement user_roles =====
	user_rolesRepo := _user_rolesRepo.NewUserRolesRepository(db)
	user_rolesService := _user_rolesService.NewUserRolesService(user_rolesRepo)
	_user_rolesHttp.NewUserRolesHttpHandler(pathPrefix, user_rolesService)
	//=====implement users =====
	usersRepo := _usersRepo.NewUsersRepository(db)
	usersService := _usersService.NewUsersService(usersRepo)
	_usersHttp.NewUsersHttpHandler(pathPrefix, usersService)
	//=====implement variant_options =====
	variant_optionsRepo := _variant_optionsRepo.NewVariantOptionsRepository(db)
	variant_optionsService := _variant_optionsService.NewVariantOptionsService(variant_optionsRepo)
	_variant_optionsHttp.NewVariantOptionsHttpHandler(pathPrefix, variant_optionsService)
	//=====implement variant_values =====
	variant_valuesRepo := _variant_valuesRepo.NewVariantValuesRepository(db)
	variant_valuesService := _variant_valuesService.NewVariantValuesService(variant_valuesRepo)
	_variant_valuesHttp.NewVariantValuesHttpHandler(pathPrefix, variant_valuesService)
	//=====implement variants =====
	variantsRepo := _variantsRepo.NewVariantsRepository(db)
	variantsService := _variantsService.NewVariantsService(variantsRepo)
	_variantsHttp.NewVariantsHttpHandler(pathPrefix, variantsService)
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
