package repository

import (
	"fmt"
	"go-search/hw16/pkg/model"
	"golang.org/x/net/context"
	"testing"
)

type App struct {
	storage Interface
	ctx     context.Context
}

var app *App

func TestMain(m *testing.M) {
	pool, err := NewPostgresDB(Config{
		User:     "postgres",
		Password: "postgres",
		Url:      "localhost",
		Port:     "5432",
		DBName:   "postgres",
	})
	if err != nil {
		fmt.Printf("Error to create DB: %v\n", err)
		return
	}

	app = &App{
		ctx:     context.Background(),
		storage: NewFilmRepository(pool),
	}

	err = app.storage.DeleteAll(app.ctx)
	if err != nil {
		fmt.Printf("Error to delete all films from TestMain: %v\n", err)
	}

	m.Run()
}

func TestDB_AddFilms(t *testing.T) {
	tests := []struct {
		name    string
		args    []model.Film
		wantErr bool
	}{
		{
			name: "Add two films",
			args: []model.Film{
				{
					Title:     "The Shawshank Redemption",
					Year:      1994,
					BoxOffice: 318412871,
					StudioId:  0,
					Rating:    "PG-13",
				},
				{
					Title:     "The Godfather",
					Year:      1972,
					BoxOffice: 91040991,
					StudioId:  1,
					Rating:    "PG-13",
				},
			},
			wantErr: false,
		},
		{
			name: "Error",
			args: []model.Film{
				{
					Title:     "The Shawshank Redemption",
					Year:      1994,
					BoxOffice: 318412871,
					StudioId:  1,
					Rating:    "wrong",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := app.storage.AddFilms(app.ctx, tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("DB.AddFilms() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Films(t *testing.T) {
	createFilmsIfNotExist(t, mockFilms)

	// Check for added films
	films, err := app.storage.Films(app.ctx, 0)
	if err != nil {
		t.Errorf("Error to get films: %v\n", err)
	}

	// tests
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Get all films",
			args: 0,
			want: 2,
		},
		{
			name: "Get second film",
			args: films[1].StudioId,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var films []model.Film

			films, err = app.storage.Films(app.ctx, tt.args)

			if err != nil {
				t.Errorf("Error to get films: %v\n", err)
			}

			if len(films) != tt.want {
				t.Errorf("Films list is wrong\n")
			}
		})
	}
}

func TestDB_UpdateFilm(t *testing.T) {
	createFilmsIfNotExist(t, mockFilms)

	films, err := app.storage.Films(app.ctx, 0)
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		args    model.Film
		wantErr bool
	}{
		{
			name: "Update film",
			args: model.Film{
				Id:     films[0].Id,
				Rating: "PG-13",
			},
			wantErr: false,
		},
		{
			name: "Error",
			args: model.Film{
				Id:     films[0].Id,
				Rating: "wrong",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := app.storage.UpdateFilm(app.ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateFilm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_DeleteAll(t *testing.T) {
	createFilmsIfNotExist(t, mockFilms)

	// tests
	err := app.storage.DeleteAll(app.ctx)
	if err != nil {
		t.Errorf("Error to delete all films: %v\n", err)
	}
}

func TestDB_DeleteFilmById(t *testing.T) {
	createFilmsIfNotExist(t, mockFilms)

	films, err := app.storage.Films(app.ctx, 0)
	if err != nil {
		t.Errorf("Error to get films: %v\n", err)
	}

	// tests
	err = app.storage.DeleteFilmById(app.ctx, films[0].Id)
	if err != nil {
		t.Errorf("Error to DeleteFilmById: %v\n", err)
	}

	films, err = app.storage.Films(app.ctx, 0)
	if err != nil {
		t.Errorf("Error to get films: %v\n", err)
	}

	if len(films) != 1 {
		t.Errorf("Films list is wrong\n")
	}
}

func addFilmsHelper(t *testing.T, films []model.Film) {
	t.Helper()
	err := app.storage.AddFilms(app.ctx, films)
	if err != nil {
		t.Errorf("Error to add films: %v\n", err)
	}
}

func createFilmsIfNotExist(t *testing.T, films []model.Film) {
	t.Helper()

	// Check for empty list
	films, err := app.storage.Films(app.ctx, 0)
	if err != nil {
		t.Errorf("Error to get films: %v\n", err)
	}

	if len(films) == 0 {
		addFilmsHelper(t, []model.Film{
			{
				Title:     "asdasd",
				Year:      1994,
				BoxOffice: 318412871,
				StudioId:  0,
				Rating:    "PG-13",
			},
			{
				Title:     "dfgh4rf",
				Year:      1972,
				BoxOffice: 91040991,
				StudioId:  1,
				Rating:    "PG-13",
			},
		})
	}

	// Check for added films
	films, err = app.storage.Films(app.ctx, 0)
	if err != nil {
		t.Errorf("Error to get films: %v\n", err)
	}

	if len(films) == 0 {
		t.Errorf("Films list is empty\n")
	}
}

var mockFilms = []model.Film{
	{Title: "The Shawshank Redemption", Year: 1994, BoxOffice: 27879069, StudioId: 1, Rating: "PG-13"},
	{Title: "The Godfather", Year: 1972, BoxOffice: 91066988, StudioId: 0, Rating: "PG-13"},
}
