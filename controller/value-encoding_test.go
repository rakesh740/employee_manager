package controller

import (
	"employee_manager/data"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_decodeEmployee(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		want    data.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeEmployee(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}
