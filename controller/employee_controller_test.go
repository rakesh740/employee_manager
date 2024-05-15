package controller

import (
	"bytes"
	"employee_manager/data"
	"employee_manager/service"
	mocks "employee_manager/service/mocks"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestNewEmployeeController(t *testing.T) {
	type args struct {
		service service.EmployeeService
	}
	tests := []struct {
		name string
		args args
		want *EmployeeController
	}{
		{name: "init test", args: args{service: nil}, want: &EmployeeController{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeController(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func TestEmployeeController_Create(t *testing.T) {

	c := GetTestGinContext()
	e := data.Employee{
		ID:       1,
		Name:     "Rakesh",
		Position: "Engineer",
		Salary:   3444440,
	}
	bEmployee, _ := json.Marshal(e)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bEmployee))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	employeeService := mocks.NewMockEmployeeService(mockCtrl)
	employeeService.EXPECT().Create(e)

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		controller *EmployeeController
		args       args
	}{
		{"create", NewEmployeeController(employeeService), args{c}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.Create(tt.args.ctx)
		})
	}
}

func TestEmployeeController_Update(t *testing.T) {

	c := GetTestGinContext()
	e := data.Employee{
		ID:       1,
		Name:     "Rakesh",
		Position: "Engineer",
		Salary:   3444440,
	}
	bEmployee, _ := json.Marshal(e)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bEmployee))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	employeeService := mocks.NewMockEmployeeService(mockCtrl)
	employeeService.EXPECT().Update(e)

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		controller *EmployeeController
		args       args
	}{
		{
			name:       "Update",
			controller: &EmployeeController{employeeService},
			args:       args{c},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.Update(tt.args.ctx)
		})
	}
}

func TestEmployeeController_Delete(t *testing.T) {
	c := GetTestGinContext()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	employeeService := mocks.NewMockEmployeeService(mockCtrl)
	employeeService.EXPECT().Delete(1)
	c.AddParam("employeeId", "1")

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		controller *EmployeeController
		args       args
	}{
		{
			name:       "Delete",
			controller: &EmployeeController{employeeService},
			args:       args{c},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.Delete(tt.args.ctx)
		})
	}
}

func TestEmployeeController_FindById(t *testing.T) {
	c := GetTestGinContext()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	employeeService := mocks.NewMockEmployeeService(mockCtrl)
	employeeService.EXPECT().FindById(1)
	c.AddParam("employeeId", "1")

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		controller *EmployeeController
		args       args
	}{
		{
			name:       "FindById",
			controller: &EmployeeController{employeeService},
			args:       args{c},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.FindById(tt.args.ctx)
		})
	}
}

func TestEmployeeController_FindAll(t *testing.T) {
	c := GetTestGinContext()
	emps := data.AllEmployeesResponse{
		Emplyoyees: []data.Employee{
			{
				ID:       1,
				Name:     "Rakesh",
				Position: "Engineer",
				Salary:   20000,
			},
			{
				ID:       2,
				Name:     "shyam",
				Position: "senior Engineer",
				Salary:   300000,
			},
			{
				ID:       3,
				Name:     "Ram",
				Position: "devops",
				Salary:   290000,
			},
		},
		Count: 3,
		Page:  1,
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	employeeService := mocks.NewMockEmployeeService(mockCtrl)
	employeeService.EXPECT().FindAll(5, 1).Return(emps)
	c.Request, _ = http.NewRequest("GET", "?page=1&limit=5", nil)

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		controller *EmployeeController
		args       args
	}{
		{
			name:       "Find all",
			controller: &EmployeeController{employeeService},
			args:       args{c},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.FindAll(tt.args.ctx)
		})
	}
}
