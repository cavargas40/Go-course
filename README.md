# Go/Golang



* [Official Site] (https://golang.org/)
* [Tour Golang] (https://tour.golang.org/welcome/2)
* [Begin with Golang] (https://www.learn-golang.org/)
* [A step ahead] (https://golangbot.com/learn-golang-series/)
* [Documentation] (https://godoc.org/)
* [Creating Golang WebServer With Echo] (https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=1)
* [GoBooks] (https://github.com/dariubs/GoBooks)
* [Unit Test] (https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d)


## Commands
Check Version
    - `$ go version` 

Run Project
    - `$ go run <filename>.go`

Documentation
    - `$ go doc <library>` ex: time

Format document
    - `$ go fmt <filename>.go`

Build Program - generates an executable file
    - `$ go build <filename>.go`

Install go package
    - `$ go get -u <packagename>`
    - `$ go get -u github.com/gorilla/mux`

Install mgo
    - `$ go get gopkg.in/mgo.v2`

Install mbson
    - `$ go get gopkg.in/mgo.v2/bson`

Api cURL Requests
### Create Movie
        curl -X POST \
        http://localhost:8080/movie \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: ce5cc735-4adc-496c-9628-6051c0defd09' \
        -H 'cache-control: no-cache' \
        -d '{
            "name": "The Wolf of Wallstreet",
            "year": 2015,
            "director": "Unknown"
        }'

### Get Movies 
        curl -X GET \
        http://localhost:8080/movies \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: 80b83e8e-4255-49b3-b2db-71ad5efe67a6' \
        -H 'cache-control: no-cache'

### Get Movie
        curl -X GET \
        http://localhost:8080/movie/5d852f9a0d06df57f7a11d88 \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: 899d64c8-768e-4846-8cbd-8029f9ff3b48' \
        -H 'cache-control: no-cache'        

### Update Movie
        curl -X PUT \
        http://localhost:8080/movie/5d852f9a0d06df57f7a11d88 \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: 859136d9-0a0b-4278-a415-5d68b5d9c293' \
        -H 'cache-control: no-cache' \
        -d '{
            "name": "The Wooooooooooolf of Wallstreet",
            "year": 2015,
            "director": "Unknown"
        }'

### Delete Movie
        curl -X DELETE \
        http://localhost:8080/movie/5d853c050d06df57f7a11f95 \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: b7894f63-6fcf-47f2-9bdc-ccb935beb62c' \
        -H 'cache-control: no-cache'