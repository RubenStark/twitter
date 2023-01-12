package bd

import(
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/RubenStark/twitter/models"
)

//Creando Teeets

func InsertoTweet(m models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid": m.UserID,
		"mensaje": m.Mensaje,
		"fecha": m.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return string(""), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}