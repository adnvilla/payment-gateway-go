SHELL := /bin/bash

default: build

build: 
	go build -o bin/payment-gateway-go-api main.go 

run: build
	./bin/payment-gateway-go-api

test:
	go test ./... -cover

mock_gen:	
	mockery