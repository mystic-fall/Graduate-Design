package handler

import (
	"context"
	"dormitory/constant"
	"dormitory/internal/repository"
	"dormitory/internal/service"
)

func (*DormitoryService) GetBuildingList(ctx context.Context, req *service.GetBuildingListRequest) (resp *service.GetBuildingListResponse, err error) {
	resp = &service.GetBuildingListResponse{}

	BuildingListPO, err := repository.GetBuildingList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.BuildingList = convert2BuildingDTO(BuildingListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2BuildingDTO(BuildingListPOs []*repository.Building) []*service.Building {
	BuildingListDTOs := make([]*service.Building, 0)
	for _, BuildingListPO := range BuildingListPOs {
		BuildingListDTOs = append(BuildingListDTOs, &service.Building{
			ID:       BuildingListPO.ID,
			DmID:     BuildingListPO.DormitoryManagerID,
			Location: BuildingListPO.Location,
			Name:     BuildingListPO.Name,
		})
	}
	return BuildingListDTOs
}

func (*DormitoryService) CreateBuilding(ctx context.Context, req *service.CreateBuildingRequest) (resp *service.CreateBuildingResponse, err error) {
	resp = &service.CreateBuildingResponse{}

	id, err := repository.CreateBuilding(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) UpdateBuilding(ctx context.Context, req *service.UpdateBuildingRequest) (resp *service.UpdateBuildingResponse, err error) {
	resp = &service.UpdateBuildingResponse{}

	err = repository.UpdateBuilding(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) DeleteBuilding(ctx context.Context, req *service.DeleteBuildingRequest) (resp *service.DeleteBuildingResponse, err error) {
	resp = &service.DeleteBuildingResponse{}

	err = repository.DeleteBuilding(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
