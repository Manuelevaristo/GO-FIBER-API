package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Nome  string             `json:"nome,omitempty" validate:"required"`
	Ativo bool               `json:"ativo,omitempty" validate:"required"`
}
