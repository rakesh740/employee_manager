package main

import (
	"employee_manager/config"
	"employee_manager/controller"
	"employee_manager/data"
	"employee_manager/helper"
	"employee_manager/repository"
	"employee_manager/router"
	"employee_manager/service"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	db := config.DatabaseConnection()
	db.Migrator().CreateTable(&data.Employee{})

	employeesRepository := repository.NewEmployeeREpositoryImpl(db)
	employeesService := service.NewEmployeeServiceImpl(employeesRepository)
	employeeController := controller.NewEmployeeController(employeesService)
	routes := router.NewRouter(employeeController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
