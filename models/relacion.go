package models

/*Relacion modelo para grabar la relación de un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioid"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}