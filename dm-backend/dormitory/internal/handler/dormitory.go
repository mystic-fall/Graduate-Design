package handler

import (
	"context"
	"dormitory/constant"
	"dormitory/internal/repository"
	"dormitory/internal/service"
)

type DormitoryService struct {
	service.UnimplementedDormitoryServiceServer
}

func NewDormitoryService() *DormitoryService {
	return &DormitoryService{}
}

func (*DormitoryService) GetDormitoryList(ctx context.Context, req *service.GetDormitoryListRequest) (resp *service.GetDormitoryListResponse, err error) {
	resp = &service.GetDormitoryListResponse{}

	DormitoryListPO, err := repository.GetDormitoryList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.DormitoryList = convert2DormitoryDTO(DormitoryListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2DormitoryDTO(DormitoryListPOs []*repository.Dormitory) []*service.Dormitory {
	DormitoryListDTOs := make([]*service.Dormitory, 0)
	for _, DormitoryListPO := range DormitoryListPOs {
		DormitoryListDTOs = append(DormitoryListDTOs, &service.Dormitory{
			ID:          DormitoryListPO.ID,
			Sn:          DormitoryListPO.Sn,
			BuildingID:  DormitoryListPO.BuildingID,
			Floor:       DormitoryListPO.Floor,
			LivedNumber: DormitoryListPO.LivedNumber,
			MaxNumber:   DormitoryListPO.MaxNumber,
		})
	}
	return DormitoryListDTOs
}

func (*DormitoryService) CreateDormitory(ctx context.Context, req *service.CreateDormitoryRequest) (resp *service.CreateDormitoryResponse, err error) {
	resp = &service.CreateDormitoryResponse{}

	id, err := repository.CreateDormitory(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) UpdateDormitory(ctx context.Context, req *service.UpdateDormitoryRequest) (resp *service.UpdateDormitoryResponse, err error) {
	resp = &service.UpdateDormitoryResponse{}

	err = repository.UpdateDormitory(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) DeleteDormitory(ctx context.Context, req *service.DeleteDormitoryRequest) (resp *service.DeleteDormitoryResponse, err error) {
	resp = &service.DeleteDormitoryResponse{}

	err = repository.DeleteDormitory(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
