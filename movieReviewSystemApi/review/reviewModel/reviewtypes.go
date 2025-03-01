package model

type Review struct {
	ID               int64  `bson:"_id,omitempty" json:"id,omitempty"`
	UserId           int64  `bson:"userId" json:"userId"`
	BaseId           int64  `bson:"baseId" json:"baseId"`
	RootId           int64  `bson:"rootId" json:"rootId"`
	HeadId           int64  `bson:"headId" json:"headId"`
	Text             string `bson:"text" json:"text"`
	Status           int64  `bson:"status" json:"status"`
	Level            int64  `bson:"level" json:"level"`
	RootCommentCount int64  `bson:"rootCommentCount" json:"rootCommentCount"`
	// TODO: Fill your own fields
	UpdateAt int64 `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt int64 `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
