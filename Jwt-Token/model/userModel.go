package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     string             `json:"firstName" validate:"required,min=2 ,max=10"`
	LastName      string             `json:"lastName" validate:"required,min=2 ,max=10"`
	Email         string             `json:"email" validate:"required,email" `
	Password      string             `json:"password" validate:"required,min=6"`
	Phone         string             `json:"phone" validate:"required"`
	Token         string             `json:"token"`
	User_Type     string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_Token string             `json:"refresh_token"`
	Created_Date  time.Time          `json:"created_Date"`
	Updated_Date  time.Time          `json:"updated_Date"`
	User_ID       string             `json:"user_ID"`
}
