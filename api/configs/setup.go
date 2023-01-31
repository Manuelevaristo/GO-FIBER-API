package configs

//Importe as dependências necessárias.
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cria uma função ConnectDB que primeiro configure o cliente para usar o URI correto e verifique se há erros
func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	//definimos um tempo limite de 10 segundos que queríamos usar ao tentar conectar
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx) //verifique se há algum erro durante a conexão com o banco de dados e cancele a conexão se o período de conexão exceder 10 segundos.
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado ao MongoDB")
	return client
}

// instância do cliente
var DB *mongo.Client = ConnectDB()

// obtendo coleções de banco de dados
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("goAPI").Collection(collectionName)
	return collection
}
