-- Include the extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the users table
-- Don't worry about passwords for this exercise
CREATE TABLE IF NOT EXISTS users
(
    id         uuid NOT NULL DEFAULT uuid_generate_v4(),
    first_name text NOT NULL,
    last_name  text NOT NULL,
    email      text unique NOT NULL,
    experience text NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp without time zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id) NOT DEFERRABLE
);

-- Create the allergies, dietary restrictions, and favorite ingredients tables
-- All of these should reference the `id` column from the users table
-- We should not have duplicate entries in any of these tables
CREATE TABLE IF NOT EXISTS allergies
(
    user_id uuid NOT NULL,
    allergy text NOT NULL,
    PRIMARY KEY (user_id, allergy) NOT DEFERRABLE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS dietary_restrictions
(
    user_id             uuid NOT NULL,
    dietary_restriction text NOT NULL,
    PRIMARY KEY (user_id, dietary_restriction) NOT DEFERRABLE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS favorite_ingredients
(
    user_id    uuid NOT NULL,
    ingredient text NOT NULL,
    PRIMARY KEY (user_id, ingredient) NOT DEFERRABLE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Set up some test data
INSERT INTO users (id, first_name, last_name, email, experience, created_at, updated_at)
VALUES ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'harry', 'potter', 'harry@hogwarts.edu', 'beginner', now(), now()),
       ('1bf5f7ba-be25-4873-b602-627e0c2e36c4', 'ash', 'ketchum', 'ash@oaklabs.org', 'pro', now(), now());

INSERT INTO allergies (user_id, allergy)
VALUES ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'tree nut'),
       ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'dairy'),
       ('1bf5f7ba-be25-4873-b602-627e0c2e36c4', 'dairy');

INSERT INTO dietary_restrictions (user_id, dietary_restriction)
VALUES ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'lactose free'),
       ('1bf5f7ba-be25-4873-b602-627e0c2e36c4', 'gluten free');

INSERT INTO favorite_ingredients (user_id, ingredient)
VALUES ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'onion'),
       ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'rutabaga'),
       ('9a784f72-b8a9-4460-a630-0b61234c87bb', 'oat milk'),
       ('1bf5f7ba-be25-4873-b602-627e0c2e36c4', 'chicken'),
       ('1bf5f7ba-be25-4873-b602-627e0c2e36c4', 'blueberry');
