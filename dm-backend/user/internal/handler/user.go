package handler

import (
	"context"
	"user/constant"
	"user/internal/repository"
	"user/internal/service"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.LoginRequest) (resp *service.LoginResponse, err error) {
	resp = &service.LoginResponse{}

	tp := req.GetType()
	switch constant.UserType(tp) {
	case constant.ADMIN:
		if pass, userID, err := repository.CheckAdmin(req); !pass {
			bizErr := err.(*constant.BizError)
			resp.UserID = userID
			resp.Code = bizErr.Code
			resp.Msg = bizErr.Msg
		} else {
			resp.Code = constant.SUCCESS_CODE
			resp.Msg = constant.SUCCESS_MSG
		}
	case constant.DM:
		if pass, userID, err := repository.CheckDM(req); !pass {
			bizErr := err.(*constant.BizError)
			resp.UserID = userID
			resp.Code = bizErr.Code
			resp.Msg = bizErr.Msg
		} else {
			resp.Code = constant.SUCCESS_CODE
			resp.Msg = constant.SUCCESS_MSG
		}
	case constant.STU:
		if pass, userID, err := repository.CheckStu(req); !pass {
			bizErr := err.(*constant.BizError)
			resp.UserID = userID
			resp.Code = bizErr.Code
			resp.Msg = bizErr.Msg
		} else {
			resp.Code = constant.SUCCESS_CODE
			resp.Msg = constant.SUCCESS_MSG
		}
	}
	return resp, nil
}
