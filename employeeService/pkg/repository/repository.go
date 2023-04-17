package repository

import (
	employeeservice "employeeService"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository interface {
	AddEmployee(employee *employeeservice.Employee) (primitive.ObjectID, error)
	DeleteEmployee(employeeID string) error
	GetEmployeesByCompany(companyID int) ([]*employeeservice.Employee, error)
	GetEmployeesByDepartment(departmentName string) ([]*employeeservice.Employee, error)
	UpdateEmployee(employeeID string, updates map[string]interface{}) error
}

type Repository struct {
	EmployeeRepository
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		EmployeeRepository: NewMongoDBRepository(db),
	}
}
