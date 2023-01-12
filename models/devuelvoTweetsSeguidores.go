package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId            string             `bson:"userid" json:"userid,omitempty"`
	UsuarioRelacionId string             `bson:"usuariorelacionid" json:"usuariorelacionid,omitempty"`
	Tweet             struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
