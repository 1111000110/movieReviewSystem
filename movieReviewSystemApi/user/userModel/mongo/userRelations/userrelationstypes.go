package model

type UserRelations struct {
	ID int64 `bson:"_id,omitempty" json:"RelationsId,omitempty"`
	// TODO: Fill your own fields
	UserId           int64 `bson:"userId" json:"userId"`
	OtherId          int64 `bson:"otherId" json:"otherId"`
	RelationshipType int64 `bson:"relationshipType" json:"relationshipType"`
	UpdateAt         int64 `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt         int64 `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
