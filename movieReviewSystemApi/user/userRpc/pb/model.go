package __

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	model "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	"strconv"
)

func ModelUserToUser(user model.User) (User, error) {
	var err error
	phone := ""
	if user.Phone != "" {
		phone, err = tool.Decrypt(user.Phone)
		if err != nil {
			return User{}, err
		}
	}
	UserId := model.GetUserIdByID(user.ID)
	return User{
		UserId:    UserId,
		Phone:     phone,
		Email:     user.Email,
		Password:  user.Password,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		Role:      user.Role,
		Status:    user.Status,
		CreateAt:  user.CreateAt,
		UpdateAt:  user.UpdateAt,
	}, nil
}
func UserToModelUser(user User) (model.User, error) {
	var err error
	phone := ""
	if user.Phone != "" {
		phone, err = tool.Encrypt(user.Phone)
		if err != nil {
			return model.User{}, err
		}
	}
	var modelUserId primitive.ObjectID
	if user.UserId == 0 {
		modelUserId, err = primitive.ObjectIDFromHex(strconv.FormatInt(user.UserId, 10))
		if err != nil {
			return model.User{}, err
		}
	}
	return model.User{
		ID:        modelUserId,
		Phone:     phone,
		Email:     user.Email,
		Password:  user.Password,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		Role:      user.Role,
		Status:    user.Status,
		CreateAt:  user.CreateAt,
		UpdateAt:  user.UpdateAt,
	}, nil
}
