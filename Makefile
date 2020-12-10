.PHONY mocs:
mocks:
	$(GOPATH)/bin/./minimock -i ./internal/app/handlers.storage -o ./internal/app/handlers/mocks/storage_mock.go
	$(GOPATH)/bin/./minimock -i ./internal/app/handlers.tgBot -o ./internal/app/handlers/mocks/tg_bot_mock.go

.PHONY test:
test:
	go test ./...