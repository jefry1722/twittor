package bd

import (
	"github.com/jefry1722/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string)(models.Usuario,bool){
	user,encontrado,_:=ChequeoYaExisteUsuario(email)
	if !encontrado{
		return user,false
	}

	passwordBytes:=[]byte(password)
	passwordBD:=[]byte(user.Password)

	err:=bcrypt.CompareHashAndPassword(passwordBD,passwordBytes)
	if err!=nil{
		return user,false
	}

	return user,true
}