package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt           int64    `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt           int64    `bson:"createAt,omitempty" json:"createAt,omitempty"`
	MovieInformationId int64    `bson:"movieInformationId,omitempty" json:"movieInformationId,omitempty"`
	Title              string   `bson:"title,omitempty" json:"title,omitempty"`
	Desc               string   `bson:"desc,omitempty" json:"desc,omitempty"`
	Author             string   `bson:"author,omitempty" json:"author,omitempty"`
	Actors             []string `bson:"actors,omitempty" json:"actors,omitempty"`
	Language           string   `bson:"language,omitempty" json:"language,omitempty"`
	Duration           int      `bson:"duration,omitempty" json:"duration,omitempty"`
	ReleaseDate        string   `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
	Genre              []string `bson:"genre,omitempty" json:"genre,omitempty"`
	Poster             string   `bson:"poster,omitempty" json:"poster,omitempty"`
}
