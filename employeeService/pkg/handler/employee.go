package handler

import (
	employeeservice "employeeService"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type addEmployeeResponse struct {
	ID string `json:"id"`
}

func (h *Handler) addEmployee(c *gin.Context) {
	var employee employeeservice.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	employeeID, err := h.services.AddEmployee(&employee)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, addEmployeeResponse{
		ID: employeeID.Hex(),
	})
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")

	err := h.services.EmployeeRepository.DeleteEmployee(employeeID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}

func (h *Handler) getEmployeesByCompany(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	employees, err := h.services.EmployeeRepository.GetEmployeesByCompany(companyID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getEmployeesByDepartment(c *gin.Context) {
	departmentName := c.Param("departmentName")

	employees, err := h.services.EmployeeRepository.GetEmployeesByDepartment(departmentName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	employeeID := c.Param("id")

	var employeeUpdates employeeservice.Employee
	err := c.BindJSON(&employeeUpdates)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updates := make(map[string]interface{})
	employeeUpdatesValue := reflect.ValueOf(employeeUpdates)
	employeeUpdatesType := reflect.TypeOf(employeeUpdates)
	for i := 0; i < employeeUpdatesValue.NumField(); i++ {
		field := employeeUpdatesType.Field(i)
		value := employeeUpdatesValue.Field(i)
		if value.Kind() == reflect.Ptr && !value.IsNil() {
			updates[field.Tag.Get("json")] = value.Elem().Interface()
		}
	}

	err = h.services.EmployeeRepository.UpdateEmployee(employeeID, updates)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}
