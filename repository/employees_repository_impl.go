package repository

import (
	"employee_manager/data"
	"employee_manager/helper"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	Db *gorm.DB
}

func NewEmployeeREpositoryImpl(Db *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{Db: Db}
}

func (t *EmployeeRepositoryImpl) Delete(employeeId int) {
	var employees data.Employee
	result := t.Db.Where("id = ?", employeeId).Delete(&employees)
	helper.ErrorPanic(result.Error)
}

func (t *EmployeeRepositoryImpl) FindAll() []data.Employee {
	var employees []data.Employee
	result := t.Db.Find(&employees)
	helper.ErrorPanic(result.Error)
	return employees
}

func (t *EmployeeRepositoryImpl) FindById(employeesId int) (employee data.Employee, err error) {
	result := t.Db.Find(&employee, employeesId)
	if result != nil {
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (t *EmployeeRepositoryImpl) Save(employee data.Employee) {
	result := t.Db.Create(&employee)
	helper.ErrorPanic(result.Error)
}

// Update implements EmployeeRepository
func (t *EmployeeRepositoryImpl) Update(employee data.Employee) {
	result := t.Db.Model(&employee).Updates(employee)
	helper.ErrorPanic(result.Error)
}
