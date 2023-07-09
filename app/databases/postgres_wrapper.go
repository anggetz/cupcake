package databases

import (
	"cupcake/interface/gateways"

	"github.com/go-pg/pg/v10"
)

type PgWrapper struct {
	db *pg.DB
}

func NewPgWrapper() gateways.Database {
	return &PgWrapper{}
}

// whereClause is a string !!!
func (i *PgWrapper) Get(tableName string, dest interface{}, whereClause interface{}) error {
	err := i.db.Model(dest).Where(whereClause.(string)).Select()
	if err != nil {
		return err
	}

	return nil
}

func (i *PgWrapper) DBClientName() string {
	return "pg"
}

func (i *PgWrapper) Close() error {
	return i.db.Close()
}

func (i *PgWrapper) Connect() (gateways.Database, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		PoolSize: 50,
		Database: "testcupcake",
	})

	i.db = db

	return i, nil

}
