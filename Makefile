.PHONY: init install test build serve clean pack deploy ship run

include .env
export $(shell sed 's/=.*//' .env)

init:
	go get -u github.com/golang/dep/cmd/dep

install: init
	rm -rf ./vendor
	dep ensure

test: install
	ginkgo -r

build: 
	rm -rf ./dist
	mkdir dist
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(APP_NAME) .
	
serve: build
	cd dist && ./$(APP_NAME)

clean:
	rm ./dist/ -rf

pack:
	docker build --build-arg APP_NAME=$(APP_NAME) -t gattal/$(APP_NAME):$(TAG) .

upload:
	docker push gattal/$(APP_NAME):$(TAG)	

run: pack
	docker run --name hello -d -p $(HOST_PORT):3333 gattal/$(APP_NAME):$(TAG)

ship: init test pack upload clean