package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
	"strconv"
)

type UserRpc struct {
	Client service.UserServiceClient
}

var UserRpcRepo UserRpc

func InitUserRpcRepo(client service.UserServiceClient) {
	UserRpcRepo.Client = client
}

func (c *UserRpc) UserLogin(req *dto.UserLoginReq) *service.LoginResponse {
	tp, err := strconv.Atoi(req.Type)
	if err != nil {
		util.LogAndPanic(err)
	}
	rpcReq := &service.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		Type:     uint32(tp),
	}
	// rpc call
	rpcResp, err := c.Client.UserLogin(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
