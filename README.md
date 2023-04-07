# The Yummly Take Home Test

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
Pass: postgres
DB: yummly
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

These four tables work together to model a simple version of a Yummly user with some taste preferences,
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