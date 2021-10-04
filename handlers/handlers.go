package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jefry1722/twittor/middlew"
	"github.com/jefry1722/twittor/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, handler y pongo a escuchar mi servidor*/
func Manejadores(){
	router:=mux.NewRouter()

	router.HandleFunc("/registro",middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT:=os.Getenv("PORT")
	if PORT==""{
		PORT = "8080"
	}

	handler:=cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}