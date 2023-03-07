package database

import (
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/xsicx/highload/internal/interfaces/config"
)

var db *sqlx.DB

func ConnectToPostgres(settings config.DBConfig) error {
	var err error

	db, err = sqlx.Open("pgx", settings.DSN)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(settings.MaxOpenConnections)
	db.SetMaxIdleConns(settings.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Second * time.Duration(settings.ConnMaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(settings.ConnMaxIdleTime))

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	if db == nil {
		return nil
	}

	return db.Close()
}

func DB() *sqlx.DB {
	return db
}
