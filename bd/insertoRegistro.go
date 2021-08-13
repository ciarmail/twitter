package bd

import (
	"context"
	"time"

	"github.com/ciarmail/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro grabo en la db un usuario
string es el id del nuevo registro
bool si tuvo exito o no
error si lo hubo*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
