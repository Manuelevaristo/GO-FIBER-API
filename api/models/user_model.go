package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/* Cria uma estrutura User com as propriedades necessárias. Adicionando omitemptye validate:"required"à tag struct
para dizer ao Fiber para ignorar os campos vazios e tornar o campo obrigatório, respectivamente. */
type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Nome  string             `json:"nome,omitempty" validate:"required"`
	Ativo bool               `json:"ativo,omitempty" validate:"required"`
}
