package repository

import (
	"github.com/jackc/pgx/v4"
	"go-search/hw16/pkg/model"
	"golang.org/x/net/context"
)

func (db *DB) Films(ctx context.Context, studio_ids ...int) ([]model.Film, error) {
	rr := struct {
		rows pgx.Rows
		err  error
	}{}

	if len(studio_ids) == 1 && studio_ids[0] != 0 {
		rr.rows, rr.err = db.pool.Query(ctx, "SELECT * FROM films WHERE studio_id = $1", studio_ids[0])
	} else {
		rr.rows, rr.err = db.pool.Query(ctx, "SELECT * FROM films")
	}

	if rr.err != nil {
		return nil, rr.err
	}

	defer rr.rows.Close()

	var films []model.Film

	for rr.rows.Next() {
		var f model.Film
		err := rr.rows.Scan(&f.Id, &f.Title, &f.Year, &f.BoxOffice, &f.StudioId, &f.Rating)
		if err != nil {
			return nil, err
		}
		films = append(films, f)
	}

	err := rr.rows.Err()
	if err != nil {
		return nil, err
	}

	return films, nil
}

func (db *DB) AddFilms(ctx context.Context, films []model.Film) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var batch = &pgx.Batch{}
	for _, f := range films {
		batch.Queue("INSERT INTO films (title, year, box_office, studio_id, rating) VALUES ($1, $2, $3, $4, $5)", f.Title, f.Year, f.BoxOffice, f.StudioId, f.Rating)
	}

	res := tx.SendBatch(ctx, batch)

	err = res.Close()
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *DB) DeleteFilmById(ctx context.Context, id int) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM films WHERE id = $1", id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *DB) DeleteAll(ctx context.Context) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM films")
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *DB) UpdateFilm(ctx context.Context, film model.Film) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE films SET title = $1, year = $2, box_office = $3, studio_id = $4, rating = $5 WHERE id = $6", film.Title, film.Year, film.BoxOffice, film.StudioId, film.Rating, film.Id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
