# Hello-World API

The source code implements this microservice (A Go code test).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

1. Linux OS - Ubuntu or Fedora is preferred.
2. git
3. Golang Setup locally on OS
4. Postman for REST API testing
5. Docker & Docker Compose

### Installing

This application can be setup locally in following ways:

#### Option A
```
go get github.com/govinda-attal/hello-world
```

#### Option B (Preferred) :heavy_check_mark:
```
cd $GOPATH/src/github.com/
mkdir govinda-attal
cd govinda-attal
git clone https://github.com/govinda-attal/hello-world.git
```

### Application Development Setup

'Makefile' will be used to setup hello-world api quickly on the development workstation.

```
cd $GOPATH/src/github.com/govinda-attal/hello-world
make install # This will go get 'dep' and use it to install vendor dependencies.
```

## Running Tests

### Unit tests

Sample BDD Unit tests are implemented using ginkgo and gomega. These unit tests are written for HTTP handler only.
Two Unit tests are written for this code test.

```
cd $GOPATH/src/github.com/govinda-attal/hello-world
make test # This will run execute unit tests with ginkgo -r command
```

## Setup

### Option A: Docker Compose - Microservice and Swagger UI as docker containers (Preferred) :heavy_check_mark:

```
cd $GOPATH/src/github.com/govinda-attal/hello-world
docker-compose up -d # This will start Hello-World Microservice and Swagger-UI which will point to Hello-World microservice swagger definition.
```

Docker compose will orchestrate containers and they can be accessed from Local OS as below:
1. Microservice on :earth_asia: http://localhost:3000
2. Swagger-UI on :earth_asia: http://localhost:8080

### Option B: Microservice is run in a docker container

```
cd $GOPATH/src/github.com/govinda-attal/hello-world

make run # This will build the docker image for microservice and spin a container locally and the API is exposed on port 3000
```

### Option C: Microservice is run locally on your OS

```
cd $GOPATH/src/github.com/govinda-attal/hello-world

make serve # This will build the Microservice and run it locally and the API is exposed on port 3000 
```

### Simple tests

1. :earth_asia: http://localhost:3000/?name=X
2. :earth_asia: http://localhost:3000/?name=error

* Note: If setup locally (no docker) then use port 3333 instead of 3000

### Swagger UI to view and trial Microservice Open API

Post running command *docker-compose up -d* use browser to open :earth_asia: http://localhost:8080

## Cleanup

For containers orchestrated by Docker-compose
```
cd $GOPATH/src/github.com/govinda-attal/hello-world
docker-compose down
docker image rm gattal/hello:1.0
```

In case when Microservice running locally as a Docker container
```
docker rm hello -f
docker image rm gattal/hello:1.0

```

In case when Microservice running locally
```
Press Ctrl+C on the terminal on which microservice was running.
```

To delete the source code for this microservice run :skull: cd $GOPATH/src/github.com/ && rm -rf  govinda-attal :skull: 

## Authors

* [Govinda Attal](https://github.com/govinda-attal)

## Acknowledgments

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
* [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* [Dependecy Injection, Mocking & TDD/BDD in Go](https://www.youtube.com/watch?v=uFXfTXSSt4I)
* [Example Usage of net/http/httptest](https://golang.org/src/net/http/httptest/example_test.go)
* [Unit testing handlers with Gorilla Mux](https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters)
* [Testing HTTP Handlers](https://blog.questionable.services/article/testing-http-handlers-go/)
* [Ginkgo - A Golang BDD Testing Framework](https://onsi.github.io/ginkgo/)
* [Using Postgres Arrays with Golang](https://www.opsdash.com/blog/postgres-arrays-golang.html)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Negroni Middleware](https://github.com/urfave/negroni)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)