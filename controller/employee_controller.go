package controller

import (
	"employee_manager/data/response"
	"employee_manager/helper"
	"employee_manager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type EmployeeController struct {
	employeesService service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeesService: service,
	}
}

func (controller *EmployeeController) Create(ctx *gin.Context) {
	log.Info().Msg("create employees")

	employee, err := decodeEmployee(ctx)
	helper.ErrorPanic(err)

	controller.employeesService.Create(employee)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Created",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) Update(ctx *gin.Context) {
	log.Info().Msg("update employees")
	employee, err := decodeEmployee(ctx)
	helper.ErrorPanic(err)

	controller.employeesService.Update(employee)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete employees")
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)
	controller.employeesService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "deleted",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid employees")
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)

	Employee := controller.employeesService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   Employee,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *EmployeeController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll employees")
	Employee := controller.employeesService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   Employee,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
