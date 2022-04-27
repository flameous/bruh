.PHONY:
test:
	@go test  -count=1 -coverprofile=cov.out ./...
	@go tool cover -html=cov.out -o coverage.html
	@open coverage.html