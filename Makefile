PAYMENT_GATEWAY_POSTGRES_DBNAME?=payment_gateway_$(PAYMENT_GATEWAY_ENV)
POSTGRES_URL=postgres://$(PAYMENT_GATEWAY_POSTGRES_USER):$(PAYMENT_GATEWAY_POSTGRES_PASSWORD)@$(PAYMENT_GATEWAY_POSTGRES_HOST)/$(PAYMENT_GATEWAY_POSTGRES_DBNAME)?sslmode=disable
PACKAGES := $(shell go list ./... | grep -v '/mock')

SHELL := /bin/bash

default: build

build: 
	go build -o bin/payment-gateway-go-api main.go 

run: build
	source .env
	./bin/payment-gateway-go-api

test:
	go test ./... -cover

test2:
	go test -cover -coverprofile=coverage.out $(PACKAGES)
	gocov convert coverage.out | gocov-html > coverage.html

mock_gen:	
	mockery

mock_stripe_gen:
	go install go.uber.org/mock/mockgen@latest
	mockgen -destination=src/pkg/stripe/mock/backend.go -package=mock github.com/stripe/stripe-go/v78 Backend

local_services_up:
	source .env && docker-compose up -d; sleep 5 # to allow postgresql to start

local_services_down:
	source .env && docker-compose down --remove-orphans

setup_migrate:
	[ -d ./src/migrations ] || mkdir -p ./src/migrations
	./install-migrate.sh

migrate: setup_migrate
	migrate -database $(POSTGRES_URL) -path ./src/migrations up	

create_migration:
	migrate create -dir ./src/migrations -format 20060102150405 -ext .sql $(name)

migrate_all:
	make migrate

init: build
	POSTGRES_USER=$(PAYMENT_GATEWAY_POSTGRES_USER) POSTGRES_PASSWORD=$(PAYMENT_GATEWAY_POSTGRES_PASSWORD) ./createdb.sh
	make migrate_all

install_dependencies:
	go install github.com/matm/gocov-html/cmd/gocov-html@latest
	go install github.com/axw/gocov/gocov@latest