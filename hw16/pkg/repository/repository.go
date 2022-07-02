package repository

import (
	"go-search/hw16/pkg/model"
	"golang.org/x/net/context"
)

type Interface interface {
	Films(ctx context.Context, studio_id ...int) ([]model.Film, error)
	AddFilms(ctx context.Context, films []model.Film) error
	DeleteFilmById(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	UpdateFilm(ctx context.Context, film model.Film) error
}
