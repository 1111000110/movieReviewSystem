package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"userId,omitempty"`
	// TODO: Fill your own fields
	Phone     string `bson:"phone" json:"phone"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	Nickname  string `bson:"nickName" json:"nickName"`
	Avatar    string `bson:"avatar" json:"avatar"`
	BirthDate int64  `bson:"birthDate" json:"birthDate"`
	Role      string `bson:"role" json:"role"`
	Gender    int64  `bson:"gender" json:"gender"`
	Status    int64  `bson:"status" json:"status"`
	UpdateAt  int64  `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  int64  `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

func GetUserIdByID(ID primitive.ObjectID) int64 {
	userIdInt, err := strconv.ParseInt(ID.Hex(), 10, 64) // 第二个参数 10 表示十进制，第三个参数 64 表示结果的位数
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return userIdInt
}
