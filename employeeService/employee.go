package employeeservice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID         *primitive.ObjectID `bson:"_id" json:"id"`
	Name       *string             `bson:"name" json:"name"`
	Surname    *string             `bson:"surname" json:"surname"`
	Phone      *string             `bson:"phone" json:"phone"`
	CompanyID  *int                `bson:"companyId" json:"companyId"`
	Passport   *Passport           `bson:"passport" json:"passport"`
	Department *Department         `bson:"department" json:"department"`
}

type Passport struct {
	Type   *string `bson:"type" json:"type"`
	Number *string `bson:"number" json:"number"`
}

type Department struct {
	Name  *string `bson:"name" json:"name"`
	Phone *string `bson:"phone" json:"phone"`
}
