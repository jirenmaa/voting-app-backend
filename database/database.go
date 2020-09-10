package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB
var ctx = context.Background()

type ConnectionOptions struct {
}

func GetConnection() *pg.DB {
	return db
}

func CreateConnection(options *pg.Options) {
	db = pg.Connect(options)
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
