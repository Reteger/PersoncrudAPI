package database

import (
	"context"
	"database/sql"
	"log"
	"personcrud/internal/config"
	"personcrud/internal/models"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Connect(cfg *config.Config) (*bun.DB, error) {

	dsn := cfg.GetDSN()

	
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	_, err = db.NewCreateTable().
		Model((*models.Person)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	log.Println("База данных подключена и таблица создана")
	return db, nil
}
