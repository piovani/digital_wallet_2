GOPATH = $(shell go env GOPATH)
GOPACKGES = $(shell go list ./... | grep -v /vendor/)

help:
	@echo "Digital Wallet 2"
	@echo "https://github.com/piovani/digital_wallet_2"
	@echo "-----------------------------------------------"
	@echo "Usege:                                         "
	@echo "make help     # prints usage info              "
	@echo "make mock     # generate mocks                 "
	@echo "make test     # run tests                      "
	@echo "make cover    # run tests and generate coverage"
	@echo "make report   # run tests and generate report  "
	@echo "make migrate  # run migrate generate database  "
	@echo "make collect  # collect price coins            "
	@echo "make start    # run api project                "
	@echo "make read     # read values coins in redis     "

mock:
	~/go/bin/mockgen -source=domain/operation.go -destination=infra/mock/operation_repository.go -package=mock
	~/go/bin/mockgen -source=domain/wallet.go -destination=infra/mock/wallet_repository.go -package=mock

test:
	CGO_ENABLE=1 go test -short -race $(GOPACKAGES)

cover:
	go test ./... -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

report:
	go test ./... -v -coverprofile cover.out .
	go tool cover -html=cover.out -o cover.html
	rm cover.out
	open ./cover.html

migrate:
	go run main.go migrate

collect:
	go run main.go collect

start:
	go run main.go api

read:
	go run main.go read