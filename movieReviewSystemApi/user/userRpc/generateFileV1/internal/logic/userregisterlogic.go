package logic

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	model "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	"time"

	"movieReviewSystem/movieReviewSystemApi/user/userRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/user/userRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func generateFormattedString() string {
	random := func() string {
		rand.Seed(time.Now().UnixNano())
		result := fmt.Sprintf("%d", rand.Intn(10)) // 随机生成0-9的数字
		return result
	}
	timestampLast8Digits := func() string {
		timestamp := time.Now().Unix() // 获取当前时间戳（秒级）
		timestampStr := fmt.Sprintf("%d", timestamp)
		length := len(timestampStr)
		if length < 8 {
			return fmt.Sprintf("%08d", timestamp) // 如果不足8位，补零
		}
		return timestampStr[length-8:] // 返回最后8位
	}
	timestampPart := timestampLast8Digits() // 获取时间戳的后8位
	return fmt.Sprintf("%s%s%s%s%s%s%s", random(), timestampPart[:1], random(), timestampPart[1:5], random(), timestampPart[5:8], random())
}

func (l *UserRegisterLogic) UserRegister(in *__.UserRegisterReq) (*__.UserRegisterResp, error) {
	// todo: add your logic here and delete this line
	if in.GetPhone() == "" || in.GetPassword() == "" {
		return nil, errors.New("phone or password is empty")
	}
	_phone, err := tool.Encrypt(in.GetPhone())
	if err != nil {
		return nil, err
	}
	_password, err := tool.HashPassword(in.GetPassword())
	if err != nil {
		return nil, err
	}
	userId, err := primitive.ObjectIDFromHex(generateFormattedString())
	if err != nil {
		return nil, err
	}
	insertData := &model.User{
		ID:       userId,
		Phone:    _phone,
		Password: _password,
		Role:     in.GetRole(),
		UpdateAt: time.Now().Unix(),
		CreateAt: time.Now().Unix(),
	}
	err = l.svcCtx.UsersModel.Insert(l.ctx, insertData)
	if err != nil {
		return nil, err
	}
	return &__.UserRegisterResp{User: &__.User{
		UserId: model.GetUserIdByID(insertData.ID),
		Role:   insertData.Role,
	}}, err
}
