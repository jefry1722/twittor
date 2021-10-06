package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jefry1722/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodo Lee los usuarios registrados en el sistema, si se recibe "R" en quienes
trae solo los que se relacionan conmigo*/
func LeoUsuariosTodo(ID string, page int64, search string, tipo string)([]*models.Usuario,bool){
	ctx,cancel:=context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db:=MongoCN.Database("Twittor")
	col:=db.Collection("usuarios")

	var results []*models.Usuario

	findOptions:=options.Find()
	findOptions.SetSkip((page-1)*20)
	findOptions.SetLimit(20)

	query:=bson.M{
		"nombre":bson.M{"$regex":`(?i)`+search},
	}

	cursor,err:=col.Find(ctx,query,findOptions)
	if err!=nil{
		fmt.Println(err.Error())
		return results,false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx){
		var s models.Usuario
		err:=cursor.Decode(&s)
		if err!=nil{
			fmt.Println(err.Error())
			return results,false
		}

		var r models.Relacion
		r.UsuarioID=ID
		r.UsuarioRelacionID=s.ID.Hex()

		incluir=false
		
		encontrado,_=ConsultoRelacion(r)
		if tipo=="new" && !encontrado{
			incluir=true
		}
		if tipo=="follow" && encontrado{
			incluir=true
		}

		if r.UsuarioRelacionID==ID{
			incluir=false
		}

		if incluir{
			s.Password=""
			s.Biografia=""
			s.SitioWeb=""
			s.Ubicacion=""
			s.Banner=""
			s.Email=""

			results=append(results, &s)
		}
	}
	err=cursor.Err()
	if err!=nil{
		fmt.Println(err.Error())
		return results,false
	}
	cursor.Close(ctx)
	return results,true
	
}
