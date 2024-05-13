package controller

import (
	"employee_manager/data"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func decodeEmployee(c *gin.Context) (data.Employee, error) {
	e := data.Employee{}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return e, err
	}
	err = json.Unmarshal(jsonData, &e)
	if err != nil {
		return e, err
	}

	return e, nil
}
