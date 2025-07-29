APP_NAME=caas-aks-api

build:
	go build -o $(APP_NAME) ./cmd

run:
	go run ./cmd/main.go

swagger:
	swag init --parseDependency --parseInternal

clean:
	rm -f $(APP_NAME)
