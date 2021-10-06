package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/jefry1722/twittor/bd"
)

/*ObtenerBanner envía el Banner al HTTP*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request){
	ID:=r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w,"Debe enviar el parámetro ID",http.StatusBadRequest)
		return
	}
	perfil,err:=bd.BuscoPerfil(ID)
	if err!=nil{
		http.Error(w,"Usuario no encontrado",http.StatusBadRequest)
		return
	}
	openFile,err:= os.Open("uploads/banners/"+perfil.Banner)
	if err!=nil{
		http.Error(w,"Imagen no encontrada",http.StatusBadRequest)
		return
	}

	_,err=io.Copy(w,openFile)
	if err!=nil{
		http.Error(w,"Error al copiar la imagen",http.StatusBadRequest)
	}
}