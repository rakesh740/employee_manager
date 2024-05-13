package service

import (
	"employee_manager/data"
)

type EmployeeService interface {
	Create(employees data.Employee)
	Update(employees data.Employee)
	Delete(employeesId int)
	FindById(employeesId int) data.Employee
	FindAll(limit, page int) data.AllEmployeesResponse
}
