package routers

import (
	"net/http"

	"github.com/jefry1722/twittor/bd"
	"github.com/jefry1722/twittor/models"
)

/*BajaRelacion realiza el borrado de la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request){
	ID:=r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID=IDUsuario
	t.UsuarioRelacionID=ID

	status,err:=bd.BorroRelacion(t)
	
	if err!=nil{
		http.Error(w,"Ocurrió un error al intentar borrar la relación "+err.Error(),http.StatusBadRequest)
		return
	}
	if !status{
		http.Error(w,"No se ha logrado borrar la relación "+err.Error(),http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}