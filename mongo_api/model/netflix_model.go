package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NetflixModel struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movive  string             `json:"movie"`
	Watched bool               `json:"watched"`
}
