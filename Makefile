mock:
	mockgen -source=src/domain/repository/connection.go -destination=mocks/repository/connection.go -package=mocks
	mockgen -source=src/domain/entity/client.go -destination=mocks/entity/client.go -package=mocks
	
tests:
	gotests -w -all ./src/usecase/$(FILE)
	# gotests -w -all ./src/infrastructure/repository_impl/$(FILE)

unit_test:
	go test ./... -v