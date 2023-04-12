# The ****** Take Home Test

## Test

### Pre-requisites

1. Docker & Docker Compose
2. An IDE (We have no preference, use what you use!)


### Now what?

1. Build a very simple API to get chef data by ID

The request here is: `curl http://localhost:8080/chef/:id`

What we want returned is something that looks like:

```json
{
    "firstName": "Harry",
    "lastName": "Potter",
    "email": "harry@hogwarts.edu",
    "experienceLevel": "beginner",
    "dietaryRestrictions": ["lactose free"],
    "allergies": ["tree nut", "dairy"],
    "favoriteIngredients": ["oat milk", "rutabaga", "onion"]
}
```

## Solution

To run:

1. build the image for the api-sever `make docker-build`
2. run containers: `make up` or `docker-compose up -d`
3. validate api: `./test.sh`

e.g.

```
curl http://localhost:8080/healthcheck
curl http://localhost:8080/chef/9a784f72-b8a9-4460-a630-0b61234c87bb | jq
curl http://localhost:8080/chef/91bf5f7ba-be25-4873-b602-627e0c2e36c4 | jq
```

## Working Notes

### Notes

- we have a DB spun up via docker compose, so I think we should match that structure, an api server started via docker-compose
- so I'll need a binary to implemetnt the server, and then wire that up to start with the db
- ideally, to validate the test, you'd still only need to run `docker-compose up -d` and the rest takes care of itself
- on the server side,the "business logic" will be transforming the db lookups into json responses

### Areas of improvement/extension

- manage environmental variable/secrets (e.g. Vault)
- more tests, always more tests :)
- ssl, cors, auth missing
- I put all the data transforamtion logic in the application layer for the sake of simplicity, but we could JOIN the tables and let te DB handle it, or use another library rather than writing the query string by hand
- add integration tests
- smarter logica aroung db reconnection

### Links

- https://github.com/golang-standards/project-layout
- https://github.com/dhax/go-base
- https://github.com/go-swagger/go-swagger
- https://github.com/docker/awesome-compose
- https://the-stack-overflow-podcast.simplecast.com/episodes/building-an-api-is-half-the-battle-qa-with-marco-palladino-from-kong/transcript
- https://dev.to/lucasnevespereira/write-a-rest-api-in-golang-following-best-practices-pe9
- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
- https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
- https://stackoverflow.com/questions/74360635/postgres-go-docker-compose-cant-ping-database-dial-tcp-127-0-0-15432-con
- 

