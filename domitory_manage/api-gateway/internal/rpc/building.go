package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
)

type DormitoryRpc struct {
	Client service.DormitoryServiceClient
}

var DormitoryRpcRepo DormitoryRpc

func InitDormitoryRpcRepo(client service.DormitoryServiceClient) {
	DormitoryRpcRepo.Client = client
}

func (c *DormitoryRpc) GetBuildingList(req *dto.GetBuildingListReq) *service.GetBuildingListResponse {
	rpcReq := &service.GetBuildingListRequest{
		Name: req.Name,
	}
	// rpc call
	rpcResp, err := c.Client.GetBuildingList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *DormitoryRpc) CreateBuilding(req *dto.CreateBuildingReq) *service.CreateBuildingResponse {
	rpcReq := &service.CreateBuildingRequest{
		Building: &service.Building{
			DmID:     req.Building.DormitoryManagerID,
			Location: req.Building.Location,
			Name:     req.Building.Name,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateBuilding(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) UpdateBuilding(req *dto.UpdateBuildingReq) *service.UpdateBuildingResponse {
	rpcReq := &service.UpdateBuildingRequest{
		Building: &service.Building{
			ID:       req.Building.ID,
			DmID:     req.Building.DormitoryManagerID,
			Location: req.Building.Location,
			Name:     req.Building.Name,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateBuilding(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) DeleteBuilding(req *dto.DeleteBuildingReq) *service.DeleteBuildingResponse {
	rpcReq := &service.DeleteBuildingRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteBuilding(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
