package router

import (
	"employee_manager/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(employeeController *controller.EmployeeController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	employeesRouter := baseRouter.Group("/employees")
	employeesRouter.GET("", employeeController.FindAll)
	employeesRouter.GET("/:employeeId", employeeController.FindById)
	employeesRouter.POST("", employeeController.Create)
	employeesRouter.PUT("/:employeeId", employeeController.Update)
	employeesRouter.DELETE("/:employeeId", employeeController.Delete)

	return router
}
