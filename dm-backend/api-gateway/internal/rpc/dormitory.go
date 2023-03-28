package rpc

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/service"
	"api-gateway/util"
	"context"
	"strconv"
)

func (c *DormitoryRpc) GetDormitoryList(req *dto.GetDormitoryListReq) *service.GetDormitoryListResponse {
	rpcReq := &service.GetDormitoryListRequest{
		Sn: req.Sn,
	}
	// rpc call
	rpcResp, err := c.Client.GetDormitoryList(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}

func (c *DormitoryRpc) CreateDormitory(req *dto.CreateDormitoryReq) *service.CreateDormitoryResponse {
	floor, _ := strconv.Atoi(req.Dormitory.Floor)
	livedNumber, _ := strconv.Atoi(req.Dormitory.LivedNumber)
	maxNumber, _ := strconv.Atoi(req.Dormitory.MaxNumber)
	rpcReq := &service.CreateDormitoryRequest{
		Dormitory: &service.Dormitory{
			BuildingID:  req.Dormitory.BuildingID,
			Floor:       uint32(floor),
			LivedNumber: uint32(livedNumber),
			MaxNumber:   uint32(maxNumber),
			Sn:          req.Dormitory.Sn,
		},
	}
	// rpc call
	rpcResp, err := c.Client.CreateDormitory(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) UpdateDormitory(req *dto.UpdateDormitoryReq) *service.UpdateDormitoryResponse {
	floor, _ := strconv.Atoi(req.Dormitory.Floor)
	livedNumber, _ := strconv.Atoi(req.Dormitory.LivedNumber)
	maxNumber, _ := strconv.Atoi(req.Dormitory.MaxNumber)
	rpcReq := &service.UpdateDormitoryRequest{
		Dormitory: &service.Dormitory{
			ID:          req.Dormitory.ID,
			BuildingID:  req.Dormitory.BuildingID,
			Floor:       uint32(floor),
			LivedNumber: uint32(livedNumber),
			MaxNumber:   uint32(maxNumber),
			Sn:          req.Dormitory.Sn,
		},
	}
	// rpc call
	rpcResp, err := c.Client.UpdateDormitory(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
func (c *DormitoryRpc) DeleteDormitory(req *dto.DeleteDormitoryReq) *service.DeleteDormitoryResponse {
	rpcReq := &service.DeleteDormitoryRequest{
		ID: req.ID,
	}
	// rpc call
	rpcResp, err := c.Client.DeleteDormitory(context.Background(), rpcReq)
	if err != nil {
		util.LogAndPanic(err)
	}
	return rpcResp
}
