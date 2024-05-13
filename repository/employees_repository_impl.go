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

type paginate struct {
	limit int
	page  int
}

func newPaginate(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) paginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Offset(offset).
		Limit(p.limit)
}

func NewEmployeeREpositoryImpl(Db *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{Db: Db}
}

func (t *EmployeeRepositoryImpl) Delete(employeeId int) {
	var employees data.Employee
	result := t.Db.Where("id = ?", employeeId).Delete(&employees)
	helper.ErrorPanic(result.Error)
}

func (t *EmployeeRepositoryImpl) FindAll(limit, page int) data.AllEmployeesResponse {
	var employees []data.Employee

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	result := t.Db.Scopes(newPaginate(limit, page).paginatedResult).Find(&employees)
	helper.ErrorPanic(result.Error)

	return data.AllEmployeesResponse{
		Emplyoyees: employees,
		Count:      len(employees),
		Page:       page,
	}
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
