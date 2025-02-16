package logic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
	model "movieReviewSystem/movieReviewSystemApi/user/userModel/mongo/user"
	"strconv"
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
func NewUserId() (int64, error) {
	// 定义随机数生成函数
	rand.Seed(time.Now().UnixNano())
	random := func() string {
		return fmt.Sprintf("%d", rand.Intn(10)) // 随机生成0-999的三位数字
	}

	// 定义获取时间戳后8位的函数
	timestampLast8Digits := func() string {
		timestamp := time.Now().Unix() // 获取当前时间戳（秒级）
		timestampStr := fmt.Sprintf("%d", timestamp)
		length := len(timestampStr)
		if length < 8 {
			return fmt.Sprintf("%08d", timestamp) // 如果不足8位，补零
		}
		return timestampStr[length-8:] // 返回最后8位
	}

	// 获取时间戳的后8位
	timestampPart := timestampLast8Digits()

	// 按照指定格式拼接字符串
	UserIdString := fmt.Sprintf("%s%s%s%s%s%s%s", random(), timestampPart[:1], random(), timestampPart[1:5], random(), timestampPart[5:8], random())
	return strconv.ParseInt(UserIdString, 10, 64)
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
	userId, err := NewUserId()
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
		UserId: insertData.ID,
		Role:   insertData.Role,
	}}, err
}
