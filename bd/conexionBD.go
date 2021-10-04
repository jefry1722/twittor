package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD*/
var MongoCN=conectarBD()
var clientOptions=options.Client().ApplyURI("mongodb+srv://jefry:admin@cluster0.acwyj.mongodb.net/Twittor")

/*ConectarBD permite conectar la BD*/
func conectarBD() *mongo.Client{
	client,err:=mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		log.Fatal(err.Error())
		return client
	}
	err=client.Ping(context.TODO(),nil)
	if err!=nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Contexión Exitosa de la BD")
	return client
}

/*ChequeoConexion es el ping a la BD*/
func ChequeoConexion() bool{
	err:=MongoCN.Ping(context.TODO(),nil)
	if err!=nil{
		return false
	}
	return true
}