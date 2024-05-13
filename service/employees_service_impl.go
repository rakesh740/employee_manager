package service

import (
	"employee_manager/data"
	"employee_manager/helper"
	"employee_manager/repository"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewEmployeeServiceImpl(employeeRepository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
	}
}

// Create implements EmployeeService
func (t *EmployeeServiceImpl) Create(employees data.Employee) {
	t.EmployeeRepository.Save(employees)
}

// Delete implements EmployeeService
func (t *EmployeeServiceImpl) Delete(employeesId int) {
	t.EmployeeRepository.Delete(employeesId)
}

// FindAll implements EmployeeService
func (t *EmployeeServiceImpl) FindAll(limit, page int) data.AllEmployeesResponse {
	result := t.EmployeeRepository.FindAll(limit, page)
	return result
}

// FindById implements EmployeeService
func (t *EmployeeServiceImpl) FindById(employeesId int) data.Employee {
	employeeData, err := t.EmployeeRepository.FindById(employeesId)
	helper.ErrorPanic(err)

	return employeeData
}

// Update implements EmployeeService
func (t *EmployeeServiceImpl) Update(employees data.Employee) {
	employeeData, err := t.EmployeeRepository.FindById(employees.ID)
	helper.ErrorPanic(err)
	employeeData.Name = employees.Name
	employeeData.Position = employees.Position
	employeeData.Salary = employees.Salary
	t.EmployeeRepository.Update(employeeData)
}
