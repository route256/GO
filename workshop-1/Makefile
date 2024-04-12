build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build
	cd notifications && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	docker-compose up --force-recreate --build

precommit:
	cd cart && make precommit
	cd loms && make precommit
	cd notifications && make precommit
