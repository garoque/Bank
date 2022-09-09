run:
	go run main.go

mock: 
	rm -rf ./mocks

	mockgen -source=./store/user/user.go -destination=./mocks/user_mock.go -package=mocks -mock_names=Store=MockUserStore
	mockgen -source=./store/transaction/transaction.go -destination=./mocks/transaction_mock.go -package=mocks -mock_names=Store=MockTRansactionStore

tests:
	@go clean -testcache
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'