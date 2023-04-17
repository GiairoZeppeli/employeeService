package service

import (
	employeeservice "employeeService"
	"employeeService/pkg/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeRepositoryService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeRepositoryService(repo repository.EmployeeRepository) *EmployeeRepositoryService {
	return &EmployeeRepositoryService{repo: repo}
}

func (s *EmployeeRepositoryService) AddEmployee(employee *employeeservice.Employee) (primitive.ObjectID, error) {
	return s.repo.AddEmployee(employee)
}

func (s *EmployeeRepositoryService) DeleteEmployee(employeeID string) error {
	return s.repo.DeleteEmployee(employeeID)
}

func (s *EmployeeRepositoryService) GetEmployeesByCompany(companyID int) ([]*employeeservice.Employee, error) {
	return s.repo.GetEmployeesByCompany(companyID)
}

func (s *EmployeeRepositoryService) GetEmployeesByDepartment(departmentName string) ([]*employeeservice.Employee, error) {
	return s.repo.GetEmployeesByDepartment(departmentName)
}

func (s *EmployeeRepositoryService) UpdateEmployee(employeeID string, updates map[string]interface{}) error {
	return s.repo.UpdateEmployee(employeeID, updates)
}
