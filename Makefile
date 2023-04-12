
REPO = api-server
BUILD_NUM = $(shell date +%s)
VERSION = v1.0.${BUILD_NUM}

up:
	docker compose --file docker-compose.yaml up --detach --remove-orphans
down:
	docker compose --file docker-compose.yaml down --volumes

go-build:
	go build -o api-server -ldflags "-X main.version=${VERSION}" -v main.go 

go-run:
	go run -ldflags "-X main.version=${VERSION}" main.go

go-test:
	go test -v ./...

docker-build:
	docker build . --file Dockerfile --build-arg BUILD="${VERSION}" --tag ${REPO}:${VERSION} --tag ${REPO}:latest

docker-run:
	docker run --interactive --tty ${REPO}:latest

list:
	docker container ls
	docker context ls
	docker image ls

clean:
	rm api-server
