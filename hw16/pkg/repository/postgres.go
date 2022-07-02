package repository

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"strconv"
)

type Config struct {
	User     string
	Password string
	Url      string
	Port     string
	DBName   string
}

type DB struct {
	pool *pgxpool.Pool
}

func NewPostgresDB(cfg Config) (*DB, error) {
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		return nil, err
	}

	connectionString := "postgresql://" + cfg.User + ":" + cfg.Password + "@" + cfg.Url + ":" + strconv.Itoa(port) + "/" + cfg.DBName + ""
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not parse connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.TODO(), config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("database ping failed: %v", err)
	}

	return &DB{pool: pool}, err
}
