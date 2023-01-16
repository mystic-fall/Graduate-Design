package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
)

func (c *UserRpc) GetStuList(req *dto.GetStuListReq) *service.GetStuListResponse {
	rpcReq := &service.GetStuListRequest{
		Sn:   req.Sn,
		Name: req.Name,
	}
	// rpc call
	rpcResp, err := c.Client.GetStuList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *UserRpc) CreateStu(req *dto.CreateStuReq) *service.CreateStuResponse {
	rpcReq := &service.CreateStuRequest{
		Student: &service.Stu{
			Sn:       req.Student.Sn,
			Sex:      req.Student.Sex,
			Name:     req.Student.Name,
			Password: req.Student.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateStu(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) UpdateStu(req *dto.UpdateStuReq) *service.UpdateStuResponse {
	rpcReq := &service.UpdateStuRequest{
		Student: &service.Stu{
			ID:       req.Student.ID,
			Sn:       req.Student.Sn,
			Sex:      req.Student.Sex,
			Name:     req.Student.Name,
			Password: req.Student.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateStu(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) DeleteStu(req *dto.DeleteStuReq) *service.DeleteStuResponse {
	rpcReq := &service.DeleteStuRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteStu(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
