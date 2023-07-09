package authentication

import (
	"cupcake/entities"
	"cupcake/interface/gateways"
	"fmt"
)

func UseCaseLogin(db *gateways.ImplementDatabase, username, password string) (bool, error) {
	// login here
	user := []entities.User{}
	err := db.Get("users", &user, []gateways.DatabaseWhereQueryBuilder{
		{
			Op:    "eq",
			Field: "name",
			Value: username,
		},
	})

	if err != nil {
		return false, err
	}

	if len(user) < 1 {
		return false, fmt.Errorf("Username or password not correct")
	}

	if user[0].Password == password {
		return true, nil
	}

	return false, fmt.Errorf("Username or password not correct")
}
