package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de conexion a la base de datos
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://adminuser:5ZwpuNDbCxGEcmVH@cluster0.mdqywer.mongodb.net/?retryWrites=true&w=majority")

// ConectarBD es la funcion que me permite conectar la base de datos
func conectarBD() *mongo.Client {
	//todo is a context that never cancels
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client

}

// ChequeoConnection es el ping a la base de datos
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
