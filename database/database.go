package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func GetConnection() *pg.DB {
	return db
}

func Init() {
	db = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT")),
		User:     os.Getenv("POSTGRESQL_USER"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
	})
}

func Ping(ctx context.Context) error {
	if err := db.Ping(ctx); err != nil {
		return err
	}

	return nil
}

func Migrate(models ...interface{}) error {
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
