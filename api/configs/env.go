package configs

//Importe as dependências necessárias.
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
	Cria uma funcão EnvMongoURI que verifique se a variável de ambiente

está carregada corretamente e retorne a variável de ambiente.
*/
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
