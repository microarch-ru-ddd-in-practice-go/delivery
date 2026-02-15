APP_NAME=delivery

.PHONY: all ci

all: generate test build
ci: all

build: test ## Build application
	mkdir -p build
	go build -o build/${APP_NAME} cmd/app/main.go

test: ## Run tests
	mockery
	go test ./...

generate: generate-server generate-grpc-clients generate-queues

generate-server:
	@rm -rf internal/generated/servers/discountsrv
	@oapi-codegen -config configs/server.cfg.yaml https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/delivery/contracts/openapi.yml

generate-grpc-clients:
	@rm -rf internal/generated/clients/*
	@curl -o ./api/proto/geo.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/geo/contracts/geo.proto
	@protoc --go_out=./internal/generated --go-grpc_out=./internal/generated ./api/proto/geo.proto

generate-queues:
	@rm -rf internal/generated/queues/*
	@curl -o ./api/proto/basket_events.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/basket/contracts/basket_events.proto
	@protoc --go_out=./internal/generated ./api/proto/basket_events.proto
	@curl -o ./api/proto/order_events.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/delivery/contracts/order_events.proto
	@protoc --go_out=./internal/generated ./api/proto/order_events.proto