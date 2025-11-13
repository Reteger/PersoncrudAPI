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
	// Используем DSN из конфигурации
	dsn := cfg.GetDSN()

	// Открываем соединение с базой данных через database/sql
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Создаем Bun DB обертку вокруг sql.DB
	db := bun.NewDB(sqldb, pgdialect.New())

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Создаём таблицу
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
