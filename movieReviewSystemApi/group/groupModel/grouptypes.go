package model

type Group struct {
	ID int64 `bson:"_id,omitempty" json:"GroupId,omitempty"`
	// TODO: Fill your own fields
	GroupName    string  `bson:"groupName" json:"GroupName"`
	GroupUrl     string  `bson:"groupUrl" json:"GroupUrl"`
	RootUserId   int64   `bson:"rootUserId" json:"RootUserId"`
	ManageUserId []int64 `bson:"manageUserId" json:"ManageUserId"`
	UserId       []int64 `bson:"userId" json:"UserId"`
	Status       int64   `bson:"status" json:"Status"`
	UpdateAt     int64   `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     int64   `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
