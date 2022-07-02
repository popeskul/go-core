package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	"go-search/hw16/pkg/model"
	"go-search/hw16/pkg/repository"
	"golang.org/x/net/context"
	"os"
)

type app struct {
	storage repository.Interface
	ctx     context.Context
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}
}

func main() {
	app := app{
		ctx: context.Background(),
	}
	var err error

	app.storage, err = repository.NewPostgresDB(repository.Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Url:      os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
	})
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Get films")
	films, err := app.storage.Films(app.ctx)
	if err != nil {
		fmt.Printf("Error to select films: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Films:", films)

	fmt.Println("Add films")
	err = app.storage.AddFilms(context.Background(), mockFilms)
	if err != nil {
		fmt.Printf("Error to add films: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Delete film")
	err = app.storage.DeleteFilmById(app.ctx, films[0].Id)
	if err != nil {
		fmt.Printf("Error to delete films: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Update film")
	newFilm := films[0]
	newFilm.Title = "UPDATED!!!!"
	err = app.storage.UpdateFilm(app.ctx, newFilm)
	if err != nil {
		fmt.Printf("Error to update films: %v\n", err)
		os.Exit(1)
	}

	// get film by id
	fmt.Println("get film by id")
	f, err := app.storage.Films(app.ctx, 0)
	if err != nil {
		fmt.Printf("Error to select films: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v", f)
}

var mockFilms = []model.Film{
	{Title: "The Shawshank Redemption", Year: 1994, BoxOffice: 27879069, StudioId: 1, Rating: "PG-13"},
	{Title: "The Godfather", Year: 1972, BoxOffice: 91066988, StudioId: 0, Rating: "PG-13"},
}
