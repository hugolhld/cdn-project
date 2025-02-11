package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	City     string             `json:"role,omitempty" validate:"required"`
}
