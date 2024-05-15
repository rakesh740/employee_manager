package router

import (
	"employee_manager/controller"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		employeeController *controller.EmployeeController
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.employeeController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
