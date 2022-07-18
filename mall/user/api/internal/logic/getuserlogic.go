package logic

import (
	"context"
	"errors"

	"go-zero-demo/mall/user/api/internal/svc"
	"go-zero-demo/mall/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	if req.Id == "" {
		return nil, errors.New("用户id不能为空")
	} else if req.Id != "1" {
		return nil, errors.New("用户" + req.Id + "不存在")
	}
	return &types.UserResp{
		Id:     req.Id,
		Name:   "头号用户",
		Gender: "保密",
	}, nil
}
