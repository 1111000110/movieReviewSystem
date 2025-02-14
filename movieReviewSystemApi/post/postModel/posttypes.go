package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	PostId   int64     `bson:"postId" json:"postId"`
	UserId   int64     `bson:"userId" json:"userId"`
	Text     string    `bson:"text" json:"text"`
	Image    []string  `bson:"image" json:"image"`
	Video    []string  `bson:"video" json:"video"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
