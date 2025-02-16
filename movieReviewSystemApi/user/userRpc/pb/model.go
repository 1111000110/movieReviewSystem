package __

import (
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	model "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	relationsModel "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/userRelations"
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
	return User{
		UserId:    user.ID,
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
	return model.User{
		ID:        user.UserId,
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
func ModelUserRelationsToUserRelations(user relationsModel.UserRelations) (Relations, error) {
	return Relations{
		RelationsId:      user.ID,
		UserId:           user.UserId,
		OtherId:          user.OtherId,
		RelationshipType: user.RelationshipType,
		CreateAt:         user.CreateAt,
		UpdateAt:         user.UpdateAt,
	}, nil
}
func UserRelationsToModelUserRelations(user Relations) (relationsModel.UserRelations, error) {
	return relationsModel.UserRelations{
		ID:               user.RelationsId,
		UserId:           user.UserId,
		OtherId:          user.OtherId,
		RelationshipType: user.RelationshipType,
		CreateAt:         user.CreateAt,
		UpdateAt:         user.UpdateAt,
	}, nil
}
