-- выборка фильмов с названием студии;
SELECT films.title, studios.title
FROM films
    JOIN studios ON films.studio_id = studios.id;
-- title            |         title
-- -----------------+-----------------------
--  Batman Begins   | Warner Bros. Pictures
--  The Dark Knight | Warner Bros. Pictures
-- (2 rows)

-- выборка фильмов для некоторого актёра;
SELECT films.title, persons.first_name, persons.last_name
FROM films
    JOIN films_actors ON films.id = films_actors.film_id
    JOIN persons ON films_actors.actor_id = persons.id
WHERE persons.first_name = 'Christian' AND persons.last_name = 'Bale';
-- title            | first_name | last_name
-- -----------------+------------+-----------
--  Batman Begins   | Christian  | Bale
--  The Dark Knight | Christian  | Bale
-- (2 rows)

-- подсчёт фильмов для некоторого режиссёра;
SELECT COUNT(*) AS films_count FROM films
    JOIN films_directors ON films.id = films_directors.film_id
    JOIN persons ON films_directors.director_id = persons.id
WHERE persons.first_name = 'Christopher' AND persons.last_name = 'Nolan';
-- films_count
-- -------------
--            2
-- (1 row)

-- выборка фильмов для нескольких режиссёров из списка (подзапрос);
SELECT films.title, persons.first_name, persons.last_name
FROM films
    JOIN films_directors ON films.id = films_directors.film_id
    JOIN persons ON films_directors.director_id = persons.id
WHERE persons.first_name IN ('Christopher', 'James') AND persons.last_name IN ('Nolan', 'Cameron');
-- title            | first_name  | last_name
-- -----------------+-------------+-----------
--  Batman Begins   | Christopher | Nolan
--  The Dark Knight | Christopher | Nolan
-- (2 rows)

-- подсчёт количества фильмов для актёра;
SELECT COUNT(*) AS films_count FROM films
    JOIN films_actors ON films.id = films_actors.film_id
    JOIN persons ON films_actors.actor_id = persons.id
WHERE persons.first_name = 'Christian' AND persons.last_name = 'Bale';
-- films_count
-- -------------
--            2
-- (1 row)

-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;
SELECT persons.first_name, persons.last_name, COUNT(*) AS films_count
FROM films
    JOIN films_actors ON films.id = films_actors.film_id
    JOIN persons ON films_actors.actor_id = persons.id
GROUP BY persons.id
HAVING COUNT(*) > 1;
-- first_name  | last_name | films_count
-- ------------+-----------+-------------
--  Christian  | Bale      |           2
-- (1 row)

-- подсчёт количества фильмов со сборами больше 1000000000;
SELECT COUNT(*) AS films_count FROM films
    WHERE films.box_office > 1000000000;
-- films_count
-- -------------
--            1
-- (1 row)

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000;
SELECT COUNT(*) AS directors_count FROM films
    JOIN films_directors ON films.id = films_directors.film_id
    JOIN persons ON films_directors.director_id = persons.id
WHERE films.box_office > 1000000000;
-- directors_count
-- ---------------
--               1
-- (1 row)

-- достать уникальные фамилии всех актёров
SELECT DISTINCT last_name FROM persons
    JOIN films_actors ON persons.id = films_actors.actor_id;

-- подсчёт количества фильмов, имеющих дубли по названию.
SELECT COUNT(*) AS films_count FROM films
    WHERE films.title IN (
        SELECT films.title
        FROM films
        GROUP BY films.title
        HAVING COUNT(*) > 1
    );
