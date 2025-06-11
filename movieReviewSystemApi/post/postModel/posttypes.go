package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	PostId   int64    `bson:"postId" json:"postId"`
	UserId   int64    `bson:"userId" json:"userId"`
	ThemeId  int64    `bson:"themeId" json:"themeId"`
	Title    string   `bson:"title" json:"title"`
	Source   int64    `bson:"source" json:"source"`
	Text     string   `bson:"text" json:"text"`
	Image    []string `bson:"image" json:"image"`
	Video    []string `bson:"video" json:"video"`
	UpdateAt int64    `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt int64    `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
