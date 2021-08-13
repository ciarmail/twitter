package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario: estructura para usuarios de la DB de MongoDB*/
type Usuario struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre     string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos  string             `bson:"apellidos,omitempty" json:"apellidos"`
	Nacimiento time.Time          `bson:"Nacimiento" json:"nacimiento,omitempty"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password,omitempty"`
	Avatar     string             `bson:"avatar" json:"avatar,omitempty"`
	Biografia  string             `bson:"biografia" json:"biografia,omitempty"`
	Banner     string             `bson:"banner" json:"banner,omitempty"`
	Ubiacion   string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb   string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
