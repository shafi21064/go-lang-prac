package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NetflixModel struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Movive  string             `json:"movie"`
	Watched bool               `json:"watched"`
}
