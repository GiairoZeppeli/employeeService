package handler

import (
	"employeeService/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/employees")
	{
		api.POST("", h.addEmployee)
		api.DELETE("/:id", h.deleteEmployee)
		api.GET("/companies/:companyId", h.getEmployeesByCompany)
		api.GET("/departments/:departmentName", h.getEmployeesByDepartment)
		api.PUT("/:id", h.updateEmployee)
	}
	return router
}
