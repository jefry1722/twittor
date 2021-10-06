package routers

import (
	"net/http"

	"github.com/jefry1722/twittor/bd"
	"github.com/jefry1722/twittor/models"
)

/*AltaRelacion realiza el reigstro de la relación entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r* http.Request){
	ID:=r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w,"El parámetro ID es obligatorio",http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID=IDUsuario
	t.UsuarioRelacionID=ID

	status,err:=bd.InsertoRelacion(t)
	if err!=nil{
		http.Error(w,"Ocurrió un error al intentar insertar la relacion"+err.Error(),http.StatusBadRequest)
		return
	}
	if !status{
		http.Error(w,"No se ha logrado insertar la relación "+err.Error(),http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}