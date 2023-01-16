package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
)

func (c *UserRpc) GetAdminList(req *dto.GetAdminListReq) *service.GetAdminListResponse {
	rpcReq := &service.GetAdminListRequest{
		Sn:   req.Sn,
		Name: req.Name,
	}
	// rpc call
	rpcResp, err := c.Client.GetAdminList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *UserRpc) CreateAdmin(req *dto.CreateAdminReq) *service.CreateAdminResponse {
	rpcReq := &service.CreateAdminRequest{
		Admin: &service.Admin{
			Name:     req.Admin.Name,
			Password: req.Admin.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateAdmin(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) UpdateAdmin(req *dto.UpdateAdminReq) *service.UpdateAdminResponse {
	rpcReq := &service.UpdateAdminRequest{
		Admin: &service.Admin{
			ID:       req.Admin.ID,
			Name:     req.Admin.Name,
			Password: req.Admin.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateAdmin(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) DeleteAdmin(req *dto.DeleteAdminReq) *service.DeleteAdminResponse {
	rpcReq := &service.DeleteAdminRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteAdmin(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
