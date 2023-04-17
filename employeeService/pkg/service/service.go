package service

import (
	employeeservice "employeeService"
	"employeeService/pkg/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeRepository interface {
	AddEmployee(employee *employeeservice.Employee) (primitive.ObjectID, error)
	DeleteEmployee(employeeID string) error
	GetEmployeesByCompany(companyID int) ([]*employeeservice.Employee, error)
	GetEmployeesByDepartment(departmentName string) ([]*employeeservice.Employee, error)
	UpdateEmployee(employeeID string, updates map[string]interface{}) error
}

type InitDb interface {
	InitDb() error
}

type Service struct {
	EmployeeRepository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EmployeeRepository: NewEmployeeRepositoryService(repos.EmployeeRepository),
	}
}
