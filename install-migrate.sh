#!/bin/bash -x

# Go 1.18+
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1
