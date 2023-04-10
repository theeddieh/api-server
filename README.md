# The ****** Take Home Test

## Test

### Pre-requisites

1. Docker & Docker Compose
2. An IDE (We have no preference, use what you use!)


### Setting up

We've provided for you a docker compose file to spin up a Postgres DB instance with some prefilled data.
You can check out the schema and data in `setup.sql`

__To start the DB:__

To start up the postgres instance, all you need to do is run `docker-compose up -d` in the 
project root directory.

The connection information:
```
User: postgres
Pass: *******
DB: ******
Schema: public
```

__To shut down the DB:__

To shut down the database run `docker compose down -v`. This will clear out the container and the volume.

### What's in the DB?

We have 4 tables and a bit of test data.

Tables:
1. `users`
2. `allergies`
3. `dietary_restrictions`
4. `favorite_ingredients`

These four tables work together to model a simple version of a ****** user with some taste preferences,
dietary restrictions, and allergies. The DB will enforce a few constraints like not allowing duplicate
allergies for a given user, but not much else.

You can find the specifics in the `setup.sql` script. Again, for convenience, we have a couple of hard-coded
user IDs that we'll be working with for this exercise. This is purely for convenience.

```
Harry: 9a784f72-b8a9-4460-a630-0b61234c87bb
Ash:   1bf5f7ba-be25-4873-b602-627e0c2e36c4
```

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

## Working Notes

### Links

- https://github.com/golang-standards/project-layout
- https://github.com/dhax/go-base
- https://github.com/go-swagger/go-swagger
- https://github.com/docker/awesome-compose
- https://the-stack-overflow-podcast.simplecast.com/episodes/building-an-api-is-half-the-battle-qa-with-marco-palladino-from-kong/transcript
- https://dev.to/lucasnevespereira/write-a-rest-api-in-golang-following-best-practices-pe9
- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
- https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
- 

### Notes

- we have a DB spun up via docker compose, so I think we should match that structure, an api server started via docker-compose
- so I'll need a binary to implemetnt the server, and then wire that up to start with the db
- ideally, to validate the test, you'd still only need to run `docker-compose up -d` and the rest takes care of itself
- 