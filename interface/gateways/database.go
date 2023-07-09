package gateways

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Database interface {
	// first is tableName
	Get(tableName string, dest interface{}, whereBuilder interface{}) error
	Close() error
	Connect() (Database, error)
	DBClientName() string
}

type DatabaseModel interface {
	TableName() string
}

type DatabaseWhereQueryBuilder struct {
	Op              string
	Field           string
	Value           interface{}
	GroupName       string
	SubQueryBuilder []DatabaseWhereQueryBuilder
}

type ImplementDatabase struct {
	database     Database
	QueryBuilder []DatabaseWhereQueryBuilder
}

func NewDatabase(implementor Database) *ImplementDatabase {
	return &ImplementDatabase{
		database: implementor,
	}
}

func (i *ImplementDatabase) Get(tableName string, dest interface{}, qBuilder []DatabaseWhereQueryBuilder) error {
	conDb, err := i.database.Connect()
	if err != nil {
		return err
	}

	defer conDb.Close()

	if conDb.DBClientName() == "pg" {

		whereClause := i.buildWhereClausePg(qBuilder, "")

		fmt.Println(whereClause)
		err = conDb.Get(tableName, dest, whereClause)
	} else if conDb.DBClientName() == "mongo" {
		whereClause := i.buildWhereClauseMongo(qBuilder, "")

		err = conDb.Get(tableName, dest, whereClause)
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
