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

// @title 	employee Service API
// @version	1.0
// @description A employee service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()

	db.Migrator().CreateTable(&data.Employee{})
	// db.Table("employees").AutoMigrate(&model.Employee{})

	// Repository
	employeesRepository := repository.NewEmployeeREpositoryImpl(db)

	// Service
	employeesService := service.NewEmployeeServiceImpl(employeesRepository)

	// Controller
	employeeController := controller.NewEmployeeController(employeesService)

	// Router
	routes := router.NewRouter(employeeController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
