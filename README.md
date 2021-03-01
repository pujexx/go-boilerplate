# go-boilerplate
Simple Golang REST API Boilerplate with clean Architecture

## To Do List
- [x] Domain generator
- [x] Repository generator
- [x] Service generator
- [x] Http Handler generator
- [x] Implement on main package generator
- [ ] Swagger generator
- [ ] Postman collection generator

## Requirment
- Database : `MySQL`
- Go v1.13.x
## Install 
Download or clone this project
```shell script
git clone https://github.com/pujexx/maheswari-boilerplate
```
Download the dependency

cd to your project 
```
cd ~/go/src/maheswari-boilerplate
```
```shell script
go mod download
```
## Run

```shell script
go run main.go
```

## Generate CRUD
to generate domain files
```shell script
go run main.go grill domain
```
to generate repository files
```shell script
go run main.go grill repository
```
to generate service files
```shell script
go run main.go grill service
```
to generate http handler files
```shell script
go run main.go grill http-handler
```
to generate implement code
```shell script
go run main.go grill implement
```
thanos gauntlet snap 
(be careful with doing this command) 
```shell script
go run main.go grill all
```

