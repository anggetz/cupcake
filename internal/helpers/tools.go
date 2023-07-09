package helpers

import "encoding/json"

func ToolCopyStruct(source interface{}, target interface{}) error {
	byteSource, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteSource, target)
	if err != nil {
		return err
	}

	return nil
}
