package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
)

func (c *DormitoryRpc) GetLiveList(req *dto.GetLiveListReq) *service.GetLiveListResponse {
	rpcReq := &service.GetLiveListRequest{}
	// rpc call
	rpcResp, err := c.Client.GetLiveList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *DormitoryRpc) CreateLive(req *dto.CreateLiveReq) *service.CreateLiveResponse {
	rpcReq := &service.CreateLiveRequest{
		Live: &service.Live{
			DormitoryID: req.Live.DormitoryID,
			StudentID:   req.Live.StudentID,
			LiveDate:    req.Live.LiveDate,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateLive(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) UpdateLive(req *dto.UpdateLiveReq) *service.UpdateLiveResponse {
	rpcReq := &service.UpdateLiveRequest{
		Live: &service.Live{
			ID:          req.Live.ID,
			DormitoryID: req.Live.DormitoryID,
			StudentID:   req.Live.StudentID,
			LiveDate:    req.Live.LiveDate,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateLive(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) DeleteLive(req *dto.DeleteLiveReq) *service.DeleteLiveResponse {
	rpcReq := &service.DeleteLiveRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteLive(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
