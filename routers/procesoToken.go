package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jefry1722/twittor/bd"
	"github.com/jefry1722/twittor/models"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string)(*models.Claim,bool,string,error){
	miClave:=[]byte("MastersDelDesarrollo")
	claims:=&models.Claim{}

	splitToken:=strings.Split(tk,"Bearer")
	if len(splitToken) !=2{
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk=strings.TrimSpace(splitToken[1])

	tkn,err:=jwt.ParseWithClaims(tk,claims,func(toke *jwt.Token)(interface{},error){
		return miClave,nil
	})

	if err==nil{
		_,encontrado,_:=bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado{
			Email=claims.Email
			IDUsuario=claims.ID.Hex()
		}
		return claims,encontrado,IDUsuario,nil
	}
	if !tkn.Valid{
		return claims,false,string(""),errors.New("token invalido")
	}
	return claims,false,string(""),err
}