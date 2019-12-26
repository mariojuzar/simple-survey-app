# Simple Survey App

Survey application that have some questions that can be answered

## Run Application

### Using go build

build application with script `go build main.go`

run application `./main`

run app with auto build `go run main.go`

### Using docker

execute this command to run app with docker

`docker-compose up --build`

note that you have to ensure DB_HOST in `.env` are the same with host in `docker-compose.yml`

shutdown docker with `docker-compose down --remove-orphans --volumes`


## Run Test

### Using Docker

execute this command to run test with docker 

`docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`

note that if you want to run test with docker, ensure that DB_HOST using the same host in `docker-compose.test.yml` and `.env` in package tests

### Using Framework Go Test

execute this command to run test with go test framework

`go test github.com/mariotj/survey-app/tests`


