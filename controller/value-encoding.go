package controller

import (
	"employee_manager/data"
	"encoding/json"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EncodeEmployee(e data.Employee) []byte {
	b, err := json.Marshal(e)
	if err != nil {
		//
	}
	return b
}

func EncodeInt(id int) []byte {
	return []byte(strconv.Itoa(id))
}

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
