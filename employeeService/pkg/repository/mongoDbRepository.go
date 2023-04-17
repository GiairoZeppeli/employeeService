package repository

import (
	"context"
	employeeservice "employeeService"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	collection *mongo.Collection
}

func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		collection: db.Collection("employees"),
	}
}

func (r *MongoDBRepository) AddEmployee(employee *employeeservice.Employee) (primitive.ObjectID, error) {
	if employee.ID.Hex() == "" {
		return primitive.NilObjectID, fmt.Errorf("employee id have nil id")
	}
	result, err := r.collection.InsertOne(context.Background(), employee)
	if err != nil {
		return primitive.NilObjectID, err
	}

	employeeID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to get employee ID")
	}

	return employeeID, nil
}

func (r *MongoDBRepository) DeleteEmployee(employeeID string) error {
	objID, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return fmt.Errorf("invalid employee ID")
	}

	filter := bson.M{"_id": objID}

	if err := employeeExistChecker(r, filter); err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %s", err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no employee found with the given ID")
	}

	return nil
}

func (r *MongoDBRepository) GetEmployeesByCompany(companyID int) ([]*employeeservice.Employee, error) {
	filter := bson.M{"companyId": companyID}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var employees []*employeeservice.Employee
	err = cur.All(context.Background(), &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil

}

func (r *MongoDBRepository) GetEmployeesByDepartment(departmentName string) ([]*employeeservice.Employee, error) {
	filter := bson.M{"department.name": departmentName}
	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var employees []*employeeservice.Employee
	err = cur.All(context.Background(), &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil

}

func (r *MongoDBRepository) UpdateEmployee(employeeID string, updates map[string]interface{}) error {
	objID, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return fmt.Errorf("invalid employee ID")
	}

	filter := bson.M{"_id": objID}

	if err := employeeExistChecker(r, filter); err != nil {
		return err
	}

	update := bson.M{"$set": updates}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update employee by ID: %s", err.Error())
	}

	return nil
}

func employeeExistChecker(repos *MongoDBRepository, filter primitive.M) error {
	var employee employeeservice.Employee
	err := repos.collection.FindOne(context.Background(), filter).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no employee found with the given ID")
		}
		return fmt.Errorf("failed to check if employee exists: %s", err.Error())
	}
	return nil
}
