package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
)

func (c *UserRpc) GetDMList(req *dto.GetDMListReq) *service.GetDMListResponse {
	rpcReq := &service.GetDMListRequest{
		Sn:   req.Sn,
		Name: req.Name,
	}
	// rpc call
	rpcResp, err := c.Client.GetDMList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *UserRpc) CreateDM(req *dto.CreateDMReq) *service.CreateDMResponse {
	rpcReq := &service.CreateDMRequest{
		DM: &service.DM{
			Sn:       req.DM.Sn,
			Sex:      req.DM.Sex,
			Name:     req.DM.Name,
			Password: req.DM.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateDM(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) UpdateDM(req *dto.UpdateDMReq) *service.UpdateDMResponse {
	rpcReq := &service.UpdateDMRequest{
		DM: &service.DM{
			ID:       req.DM.ID,
			Sn:       req.DM.Sn,
			Sex:      req.DM.Sex,
			Name:     req.DM.Name,
			Password: req.DM.Password,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateDM(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *UserRpc) DeleteDM(req *dto.DeleteDMReq) *service.DeleteDMResponse {
	rpcReq := &service.DeleteDMRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteDM(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
