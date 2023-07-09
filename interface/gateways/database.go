package gateways

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Database interface {
	// first is tableName
	Get(tableName string, dest interface{}, whereBuilder interface{}) error
	Close() error
	Connect(DatabaseOption) (Database, error)
	DBClientName() string
}

type DatabaseOption struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

type DatabaseWhereQueryBuilder struct {
	Op              string
	Field           string
	Value           interface{}
	GroupName       string
	SubQueryBuilder []DatabaseWhereQueryBuilder
}

type ImplementDatabase struct {
	database       Database
	databaseOption DatabaseOption
	QueryBuilder   []DatabaseWhereQueryBuilder
}

func NewDatabase(implementor Database, databaseOption *DatabaseOption) *ImplementDatabase {

	return &ImplementDatabase{
		database:       implementor,
		databaseOption: *databaseOption,
	}
}

func (i *ImplementDatabase) Close() error {
	return i.database.Close()
}

func (i *ImplementDatabase) Get(tableName string, dest interface{}, qBuilder []DatabaseWhereQueryBuilder) error {
	conDb, err := i.database.Connect(i.databaseOption)
	if err != nil {
		return err
	}

	if conDb.DBClientName() == "pg" {

		whereClause := i.buildWhereClausePg(qBuilder, "")

		err = conDb.Get(tableName, dest, whereClause)
	} else if conDb.DBClientName() == "mongo" {
		whereClause := i.buildWhereClauseMongo(qBuilder, "")

		err = conDb.Get(tableName, dest, whereClause)
	} else if conDb.DBClientName() == "mock" {
		err = conDb.Get(tableName, dest, "")
	} else {
		return fmt.Errorf(conDb.DBClientName(), "NOT SUPPORTTED")
	}

	return err

}

func (i *ImplementDatabase) buildWhereClauseMongo(qBuilder []DatabaseWhereQueryBuilder, groupName string) bson.M {
	whereClause := bson.M{}

	for _, qBuilder := range qBuilder {

		if qBuilder.GroupName != "" {
			whereClause = bson.M{
				"$" + qBuilder.GroupName: i.buildWhereClauseMongo(qBuilder.SubQueryBuilder, ""),
			}
		} else {
			whereClause = bson.M{
				qBuilder.Field: bson.M{
					"$" + qBuilder.Op: qBuilder.Value,
				},
			}
		}

	}

	return whereClause
}

func (i *ImplementDatabase) buildWhereClausePg(qBuilder []DatabaseWhereQueryBuilder, groupName string) string {
	whereClause := ""
	if groupName == "" {
		groupName = "and"
	}

	opTranslatorPg := make(map[string]string)
	opTranslatorPg["eq"] = "="
	whereClause += "("
	for index, qBuilder := range qBuilder {

		if index > 0 {
			whereClause += groupName
		}

		if qBuilder.GroupName != "" {
			whereClause += " and " + i.buildWhereClausePg(qBuilder.SubQueryBuilder, qBuilder.GroupName)
		} else {
			whereClause += qBuilder.Field + " " + opTranslatorPg[qBuilder.Op] + " " + qBuilder.Value.(string)
		}

	}
	whereClause += ")"

	return whereClause
}

// database query builder

func (qb *DatabaseWhereQueryBuilder) WhereAnd(subQBuilder []DatabaseWhereQueryBuilder) {
	qb.GroupName = "AND"
	qb.SubQueryBuilder = subQBuilder
}

func (qb *DatabaseWhereQueryBuilder) WhereEq(field, value string) {
	qb.Field = field
	qb.Op = "eq"
	qb.Value = value
	qb.GroupName = ""
}
