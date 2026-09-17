LOCAL_BIN := $(CURDIR)/bin
IMAGE_REPO:=gitlab-registry.ozon.dev/go/classroom-20/experts/product-service
IMAGE_TAG:=v1

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.swagger-validate:
	$(LOCAL_BIN)/swagger -q validate ./api/products.yaml

.generate:
	$(LOCAL_BIN)/swagger generate server --server-package=internal/app/http/server --model-package=internal/app/http/model -f ./api/products.yaml

generate: .swagger-validate .generate
	go mod tidy

build:
	go build -o ./bin/products ./cmd/products-server

run:
	go run ./cmd/products-server  --port 8082 

.PHONY: .bin-deps \
		.swagger-validate \
		.generate \
		generate \
		run \
		build

build-docker-image:
	docker build --platform=linux/amd64 -t $(IMAGE_REPO):$(IMAGE_TAG) .

push-docker-image: build-docker-image
	docker push $(IMAGE_REPO):$(IMAGE_TAG)
