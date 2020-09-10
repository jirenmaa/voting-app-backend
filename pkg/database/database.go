package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Database struct {
	*pg.DB
}

func NewConnection(options *pg.Options) (*Database, error) {
	db := pg.Connect(options)

	if err := PingDB(db); err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

func PingDB(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func (d *Database) CreateSchema(models ...interface{}) error {
	for _, model := range models {
		err := d.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
