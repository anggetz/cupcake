package presenters

import (
	"cupcake/entities"
	"cupcake/internal/helpers"
)

type UserAllDataView struct {
	Name string
	// Password string
}

func UserAllData(users []entities.User) ([]UserAllDataView, error) {
	// remove password because it is sensitive data
	// no implement

	targetView := []UserAllDataView{}
	err := helpers.ToolCopyStruct(users, &targetView)
	if err != nil {
		return nil, err
	}

	return targetView, nil

}
