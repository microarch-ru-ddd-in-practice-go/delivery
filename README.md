# Демо проект к курсу "Domain Driven Design и Clean Architecture на языке Go"
📚 Подробнее о курсе: [https://microarch.ru/courses/ddd/languages/go](https://microarch.ru/courses/ddd/languages/go?utm_source=gitlab&utm_medium=repository&utm_content=basket)

---

# HTTP (генерация HTTP сервера)
```
make generate-server
```

# gRPC (генерация gRPC клиента)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH="$PATH:$(go env GOPATH)/bin"

make generate-grpc-clients
```

# Kafka (генерация интеграционных сообщений)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH="$PATH:$(go env GOPATH)/bin"

make generate-queues
```

# Тестирование
```
make test
```

# Качество кода
```
golangci-lint run
```

# Документация используемых библилиотек
* [Oapi-codegen] (https://github.com/oapi-codegen/oapi-codegen)
* [Protobuf] (https://protobuf.dev/reference/go/go-generated/)
* [gRPC] (https://grpc.io/docs/languages/go/)
* [Mockery] (https://vektra.github.io/mockery/latest/)

# Лицензия

# Запросы к БД
```
SELECT * public.assignments;
SELECT * FROM public.couriers;
SELECT * FROM public.orders;
SELECT * public.outbox;
```

# Очистка БД (все кроме справочников)
```
DELETE FROM public.assignments;
DELETE FROM public.couriers;
DELETE FROM public.orders;

DELETE FROM public.outbox;
```

# Лицензия

Код распространяется под лицензией [MIT](./LICENSE).  
© 2025 microarch.ru