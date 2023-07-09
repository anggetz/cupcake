package authentication

import (
	"cupcake/app/databases"
	"cupcake/entities"
	"cupcake/interface/gateways"
	"testing"
)

func TestPositiveLogin(t *testing.T) {
	t.Run("run negative", func(t *testing.T) {
		db := gateways.NewDatabase(databases.NewMockClient([]interface{}{
			[]entities.User{
				{
					Name:     "test",
					Password: "test2",
				},
			}, //dummies for login
		}), &gateways.DatabaseOption{})
		defer db.Close()

		ret, err := UseCaseLogin(db, "test", "test")
		if err != nil {
			t.Log(err.Error())
		}

		if ret {
			t.Error("test case must be false")
		}
	})

	t.Run("run positif", func(t *testing.T) {
		db := gateways.NewDatabase(databases.NewMockClient([]interface{}{
			[]entities.User{
				{
					Name:     "test",
					Password: "test",
				},
			}, //dummies for login
		}), &gateways.DatabaseOption{})
		defer db.Close()

		ret, err := UseCaseLogin(db, "test", "test")
		if err != nil {
			t.Log(err.Error())
		}

		if !ret {
			t.Error("test case must be true")
		}
	})

	t.Run("run wrong username", func(t *testing.T) {
		db := gateways.NewDatabase(databases.NewMockClient([]interface{}{}), &gateways.DatabaseOption{})
		defer db.Close()

		ret, err := UseCaseLogin(db, "test", "test")
		if err != nil {
			t.Log(err.Error())
		}

		if ret {
			t.Error("test case must be false")
		}
	})
}
