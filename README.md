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


## API Documentation

List of all API in survey app:

1. Create Survey
    
    `POST /api/v1/survey`
    
    request body
    
    ```
   {"name" : "desired survey name"}
   ```
    
2. Update Survey

    `PUT /api/v1/survey`
    
    request body
        
    ```
   {"id": id, "name" : "desired survey name"}
   ```
    
3. Delete Survey by id

    `DELETE /api/v1/survey/:id`
    
4. Get survey by id

    `GET /api/v1/survey/:id`
    
5. Get all survey

    `GET /api/v1/survey`
    
6. Add question by survey id

    `POST /api/v1/question`
    
    request body
        
    ```
   {"question" : "desired question", "survey_id": id}
   ```
    
7. Update Question by question id

    `PUT /api/v1/question`
    
    request body
            
    ```
   {"id": id, "question" : "desired update question"}
   ```
    
8. Get all question by survey id

    `GET /api/v1/question/survey/:id`
    
9. Delete question by question id

    `DELETE /api/v1/question/question/:id`
    
10. Delete all question by survey id

    `DELETE /api/v1/question/survey/:id`
    
11. Add answer to a question

    `POST /api/v1/survey/answer`
    
    request body
            
   ```
   {"answer" : "desired answer", "username": "username", "survey_question_id": id}
   ```
    
12. Update answer by id

    `PUT /api/v1/survey/answer`
    
    request body
                
   ```
   {"answer" : "desired answer", "id": id}
   ```
    
13. Get report for survey by survey id (including question and answer)

    `GET /api/v1/report/survey/:id`