package responses

import "github.com/gofiber/fiber/v2"

//cria uma estrutura UserResponse com Status, Message e Data, propriedade para representar o tipo de resposta da API.
type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
