# workshop-2

## Генрация go файлов из proto файлов

### `protoc` 

Официальный компилятор protobuf для генерации кода из proto-файлов

1. Установка: https://grpc.io/docs/protoc-installation/
1. man pages: https://manpages.debian.org/testing/protobuf-compiler/protoc.1.en.html
1. Для генерации нам необходимо:
    - Установить необходимые плагины (см. _Makefile_ `.bin-deps`)
    - Завендорить внешние `proto` зависимости (см. _Makefile_ `.vendor-proto`)
    - Указать все пути для импорта proto-файлов, плагинов и их опции вызова (см. Makefile `.protoc-generate`)
1. Вызвать команду `make generate` (см. _Makefile_)

### `buf`

Фреймворк для генерации кода из proto файлов

1. Сайт: https://buf.build/
1. Установка: https://buf.build/docs/installation
1. Для генерации нам необходимо:
    - Установить `buf` (см. _Makefile_ `.bin-deps`)
    - Завендорить внешние `proto` зависимости (см. `buf.yaml`)
    - Указать используемые плагины и их опции вызова (см. `buf.gen.yaml`)
1. Вызвать команду `make generate-buf` (см. _Makefile_)


## Плагины

### protoc-gen-go

Документация: https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go

Плагин необходим для генерации `go` типов из `protobuf`. Сгенерированный код находится в файлах с расширением `.pb.go`.

### protoc-gen-go-grpc

Документация: https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc

Плагин необходим для генерации реализации `gRPC` клиента и сервера на `go` из `protobuf`. Сгенерированный код находится в файлах с расширением `_grpc.pb.go`.


### protoc-gen-grpc-gateway

Документация: 
- https://pkg.go.dev/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
- https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
- https://github.com/grpc-ecosystem/grpc-gateway


Плагин необходим для генерации RESTful HTTP API прокси-сервера на `go` из `protobuf`. Сгенерированный код находится в файлах с расширением `.pb.gw.go`.

### protoc-gen-openapiv2

Документация: 
- https://pkg.go.dev/github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
- https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/


Плагин необходим для генерации openapiv2 спецификации для вашего RESTful HTTP API прокси-сервера сгенерированного из `protobuf`. Спецификация находится в файлах с расширением `.swagger.json`.


### protoc-gen-validate

Документация: 
- https://pkg.go.dev/github.com/deelawn/protoc-gen-validation
- https://github.com/bufbuild/protoc-gen-validate/blob/main/docs.md


Плагин необходим для генерации функций валидации ваших `protobuf` сообщений на `go`.
