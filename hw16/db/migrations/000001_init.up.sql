DROP TABLE IF EXISTS films_directors;
DROP TABLE IF EXISTS films_actors;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS studios;
DROP TABLE IF EXISTS persons;
DROP TYPE IF EXISTS rating;

CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE studios (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE films (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    year INTEGER DEFAULT 0,
    box_office INTEGER DEFAULT 0,
    studio_id INTEGER REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    rating rating
);

CREATE TABLE films_actors (
    id BIGSERIAL PRIMARY KEY,
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE ON UPDATE CASCADE,
    actor_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE(film_id, actor_id)
);

CREATE TABLE films_directors (
    id BIGSERIAL PRIMARY KEY,
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE ON UPDATE CASCADE,
    director_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE(film_id, director_id)
);
