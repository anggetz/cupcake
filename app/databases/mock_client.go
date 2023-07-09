package databases

import (
	"cupcake/interface/gateways"
	"cupcake/internal/helpers"
)

type MockClient struct {
	dummiesValues []interface{}
	dummiesIndex  int
}

func NewMockClient(dummiesValues []interface{}) gateways.Database {
	return &MockClient{
		dummiesValues: dummiesValues,
		dummiesIndex:  0,
	}
}

func (i *MockClient) Get(tableName string, dest interface{}, matchAggregate interface{}) error {
	if i.dummiesIndex <= len(i.dummiesValues)-1 {
		err := helpers.ToolCopyStruct(i.dummiesValues[i.dummiesIndex], dest)
		if err != nil {
			return err
		}
		i.dummiesIndex++
	}

	return nil
}

func (i *MockClient) DBClientName() string {
	return "mock"
}

func (i *MockClient) Close() error {
	// closed
	return nil
}

func (i *MockClient) Connect(dbOption gateways.DatabaseOption) (gateways.Database, error) {
	// connected

	return i, nil

}
