
REPO = "theeddieh/jubilant-tribble"
VERSION = "v1"
ENV_VARS = "makefile environmental variable"

api-build:
	docker build --tag ${REPO}:${VERSION} .

api-run:
	docker run --interactive --tty --env TEST_ENV=${ENV_VARS} ${REPO}:${VERSION}

api-start:
	docker compose --file docker-compose-API.yaml up --detach --remove-orphans

api-stop:
	docker compose --file docker-compose-API.yaml down --volumes

list:
	docker container ls
	docker context ls
	docker image ls

db-start:
	docker compose --file docker-compose-DB.yaml up --detach --remove-orphans

db-stop:
	docker compose --file docker-compose-DB.yaml down --volumes


