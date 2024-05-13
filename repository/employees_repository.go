package repository

import (
	"employee_manager/data"
)

type EmployeeRepository interface {
	Save(employees data.Employee)
	Update(employees data.Employee)
	Delete(employeesId int)
	FindById(employeesId int) (employees data.Employee, err error)
	FindAll(limit, page int) data.AllEmployeesResponse
}
