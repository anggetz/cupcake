package helpers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

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

func GetGinContextBody(c *gin.Context, dest interface{}) error {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		return err
	}

	return json.Unmarshal(jsonData, dest)
}
