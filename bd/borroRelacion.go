package bd

import (
	"context"
	"time"

	"github.com/jefry1722/twittor/models"
)

/*BorroRelacion borra la relación de la BD*/
func BorroRelacion(t models.Relacion)(bool,error){
	ctx,cancel:=context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db:=MongoCN.Database("Twittor")
	col:=db.Collection("relacion")

	_,err:=col.DeleteOne(ctx,t)
	if err!=nil{
		return false,err
	}
	return true,nil
}