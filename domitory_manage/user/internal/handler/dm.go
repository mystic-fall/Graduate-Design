package handler

import (
	"context"
	"user/constant"
	"user/internal/repository"
	"user/internal/service"
)

func (*UserService) GetDMList(ctx context.Context, req *service.GetDMListRequest) (resp *service.GetDMListResponse, err error) {
	resp = &service.GetDMListResponse{}

	DMListPO, err := repository.GetDMList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.DMList = convert2DMDTO(DMListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2DMDTO(DMListPOs []*repository.DormitoryManager) []*service.DM {
	DMListDTOs := make([]*service.DM, 0)
	for _, DMListPO := range DMListPOs {
		DMListDTOs = append(DMListDTOs, &service.DM{
			ID:       DMListPO.ID,
			Sn:       DMListPO.Sn,
			Name:     DMListPO.Name,
			Sex:      DMListPO.Sex,
			Password: DMListPO.Password,
		})
	}
	return DMListDTOs
}

func (*UserService) CreateDM(ctx context.Context, req *service.CreateDMRequest) (resp *service.CreateDMResponse, err error) {
	resp = &service.CreateDMResponse{}

	id, err := repository.CreateDM(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) UpdateDM(ctx context.Context, req *service.UpdateDMRequest) (resp *service.UpdateDMResponse, err error) {
	resp = &service.UpdateDMResponse{}

	err = repository.UpdateDM(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) DeleteDM(ctx context.Context, req *service.DeleteDMRequest) (resp *service.DeleteDMResponse, err error) {
	resp = &service.DeleteDMResponse{}

	err = repository.DeleteDM(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
