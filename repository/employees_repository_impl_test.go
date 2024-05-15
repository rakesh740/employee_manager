package repository

import (
	"employee_manager/data"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseTestSuite is the test suite.
type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *DatabaseTestSuite) SetupSuite() {
	var host, port, employee, password, dbName = "localhost", 5432, "postgres", "", "testdb"
	sqlInfo := fmt.Sprintf("host=%s port=%d employee=%s password=%s dbname=%s sslmode=disable", host, port, employee, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	suite.Require().NoError(err, "Error connecting to the test database")

	// Enable logging for Gorm during tests
	suite.db = db.Debug()

	err = suite.db.AutoMigrate(&data.Employee{})
	suite.Require().NoError(err, "Error auto-migrating database tables")
}

// TestemployeeInsertion tests inserting a employee record.
func (suite *DatabaseTestSuite) TestemployeeInsertion() {
	// Create a employee
	employee := data.Employee{Name: "John Doe", Salary: 13000, Position: "Engineer"}
	err := suite.db.Create(&employee).Error
	suite.Require().NoError(err, "Error creating employee record")

	// Retrieve the inserted employee
	var retrievedemployee data.Employee
	err = suite.db.First(&retrievedemployee, "name = ?", "John Doe").Error
	suite.Require().NoError(err, "Error retrieving employee record")

	// Verify that the retrieved employee matches the inserted employee
	suite.Equal(employee.Salary, retrievedemployee.Salary, "Salary should match")
	suite.Equal(employee.Position, retrievedemployee.Position, "Position should match")
}

// TearDownSuite is called once after the test suite runs.
func (suite *DatabaseTestSuite) TearDownSuite() {
	// Clean up: Close the database connection
	err := suite.db.Exec("DROP TABLE employees;").Error
	suite.Require().NoError(err, "Error dropping test table")

	suite.Require().NoError(err, "Error closing the test database")
}

// TestSuite runs the test suite.
func TestSuite(t *testing.T) {
	// Skip the tests if the PostgreSQL connection details are not provided
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("Skipping PostgreSQL tests; provide POSTGRES_DSN environment variable.")
	}

	suite.Run(t, new(DatabaseTestSuite))
}

func Test_newPaginate(t *testing.T) {
	type args struct {
		limit int
		page  int
	}
	tests := []struct {
		name string
		args args
		want *paginate
	}{
		{"init", args{1, 10}, &paginate{limit: 1, page: 10}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPaginate(tt.args.limit, tt.args.page); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPaginate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paginate_paginatedResult(t *testing.T) {

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		p    *paginate
		args args
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.paginatedResult(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paginate.paginatedResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmployeeREpositoryImpl(t *testing.T) {
	type args struct {
		Db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want EmployeeRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeREpositoryImpl(tt.args.Db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeREpositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeRepositoryImpl_Delete(t *testing.T) {
	type args struct {
		employeeId int
	}
	tests := []struct {
		name string
		tr   *EmployeeRepositoryImpl
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Delete(tt.args.employeeId)
		})
	}
}

func TestEmployeeRepositoryImpl_FindAll(t *testing.T) {
	type args struct {
		limit int
		page  int
	}
	tests := []struct {
		name string
		tr   *EmployeeRepositoryImpl
		args args
		want data.AllEmployeesResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.FindAll(tt.args.limit, tt.args.page); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryImpl.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeRepositoryImpl_FindById(t *testing.T) {

	type args struct {
		employeesId int
	}
	tests := []struct {
		name         string
		tr           *EmployeeRepositoryImpl
		args         args
		wantEmployee data.Employee
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmployee, err := tt.tr.FindById(tt.args.employeesId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeeRepositoryImpl.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEmployee, tt.wantEmployee) {
				t.Errorf("EmployeeRepositoryImpl.FindById() = %v, want %v", gotEmployee, tt.wantEmployee)
			}
		})
	}
}

func TestEmployeeRepositoryImpl_Save(t *testing.T) {
	type args struct {
		employee data.Employee
	}
	tests := []struct {
		name string
		tr   *EmployeeRepositoryImpl
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Save(tt.args.employee)
		})
	}
}

func TestEmployeeRepositoryImpl_Update(t *testing.T) {
	type args struct {
		employee data.Employee
	}
	tests := []struct {
		name string
		tr   *EmployeeRepositoryImpl
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Update(tt.args.employee)
		})
	}
}
