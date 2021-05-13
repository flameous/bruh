#!/bin/bash

go test ./... -coverprofile=cov.out && go tool cover -html=cov.out -o coverage.html && rm cov.out && firefox coverage.html 
