bootstrap:
	@go get github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/swaggo/swag/cmd/swag

test:
	@ginkgo -r

