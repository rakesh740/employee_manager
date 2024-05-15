package service

import (
	"employee_manager/data"
	"employee_manager/repository"
	mock_repository "employee_manager/repository/mocks"

	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewEmployeeServiceImpl(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	employeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)

	type args struct {
		employeeRepository repository.EmployeeRepository
	}

	tests := []struct {
		name string
		args args
		want EmployeeService
	}{
		{
			name: "not equal check",
			args: args{employeeService},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeServiceImpl(tt.args.employeeRepository); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeServiceImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeServiceImpl_Create(t *testing.T) {
	type args struct {
		employees data.Employee
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mEmployeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)
	emp := data.Employee{
		Name:     "Rakshit",
		Position: "Engineer",
		Salary:   230000,
	}
	mEmployeeService.EXPECT().Save(emp)

	tests := []struct {
		name string
		tr   *EmployeeServiceImpl
		args args
	}{
		{
			name: "save employee",
			tr: &EmployeeServiceImpl{
				mEmployeeService,
			},
			args: args{emp},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Create(tt.args.employees)
		})
	}
}

func TestEmployeeServiceImpl_Delete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mEmployeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)
	mEmployeeService.EXPECT().Delete(1)

	type args struct {
		employeesId int
	}
	tests := []struct {
		name string
		tr   *EmployeeServiceImpl
		args args
	}{
		{
			name: "delete success",
			tr:   &EmployeeServiceImpl{mEmployeeService},
			args: args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Delete(tt.args.employeesId)
		})
	}
}

func TestEmployeeServiceImpl_FindAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mEmployeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)
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
	mEmployeeService.EXPECT().FindAll(1, 5).Return(emps)

	type args struct {
		limit int
		page  int
	}
	tests := []struct {
		name string
		tr   *EmployeeServiceImpl
		args args
		want data.AllEmployeesResponse
	}{
		{
			name: "FindAll success",
			tr:   &EmployeeServiceImpl{mEmployeeService},
			args: args{1, 5},
			want: emps,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.FindAll(tt.args.limit, tt.args.page); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServiceImpl.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeServiceImpl_FindById(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mEmployeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)
	emp := data.Employee{
		ID:       1,
		Name:     "Rakshit",
		Position: "Engineer",
		Salary:   230000,
	}
	mEmployeeService.EXPECT().FindById(1).Return(emp, nil)

	type args struct {
		employeesId int
	}
	tests := []struct {
		name string
		tr   *EmployeeServiceImpl
		args args
		want data.Employee
	}{
		{
			name: "FindById success",
			tr:   &EmployeeServiceImpl{mEmployeeService},
			args: args{1},
			want: emp,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.FindById(tt.args.employeesId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServiceImpl.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeServiceImpl_Update(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mEmployeeService := mock_repository.NewMockEmployeeRepository(mockCtrl)
	emp := data.Employee{
		ID:       1,
		Name:     "Rakshit",
		Position: "Engineer",
		Salary:   230000,
	}
	mEmployeeService.EXPECT().Update(emp)

	type args struct {
		employees data.Employee
	}
	tests := []struct {
		name string
		tr   *EmployeeServiceImpl
		args args
	}{
		{
			name: "update success",
			tr:   &EmployeeServiceImpl{mEmployeeService},
			args: args{emp},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Update(tt.args.employees)
		})
	}
}
