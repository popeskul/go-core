INSERT INTO studios (id, title) VALUES (0, 'Warner Bros. Pictures');
INSERT INTO studios (id, title) VALUES (1, 'Universal Pictures');
ALTER SEQUENCE studios_id_seq RESTART WITH 100;

-- directors
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (0, 'Christopher', 'Nolan', 1970);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (1, 'James', 'Cameron', 1970);

-- Actors
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (2, 'Christian', 'Bale', 1974);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (3, 'Cillian', 'Murphy', 1979);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (4, 'Heath', 'Ledger', 1980);

ALTER SEQUENCE persons_id_seq RESTART WITH 100;

-- Movies
INSERT INTO films (id, title, year, box_office, studio_id, rating) VALUES (0, 'Batman Begins', 2005, 373700000, 0, 'PG-13');
INSERT INTO films (id, title, year, box_office, studio_id, rating) VALUES (1, 'The Dark Knight', 2008, 1006000000, 0, 'PG-13');
ALTER SEQUENCE films_id_seq RESTART WITH 100;

-- Batman Begins
-- Bind actors to film
INSERT INTO films_actors (film_id, actor_id) VALUES (0, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (0, 3);
-- Bind director to film
INSERT INTO films_directors (film_id, director_id) VALUES (0, 0);

-- The Dark Knight
-- Bind actors to film
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 4);
-- Bind director to film
INSERT INTO films_directors (film_id, director_id) VALUES (1, 0);
