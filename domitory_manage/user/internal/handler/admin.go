package handler

import (
	"context"
	"user/constant"
	"user/internal/repository"
	"user/internal/service"
)

func (*UserService) GetAdminList(ctx context.Context, req *service.GetAdminListRequest) (resp *service.GetAdminListResponse, err error) {
	resp = &service.GetAdminListResponse{}

	adminListPO, err := repository.GetAdminList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.AdminList = convert2AdminDTO(adminListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2AdminDTO(adminListPOs []*repository.Admin) []*service.Admin {
	adminListDTOs := make([]*service.Admin, 0)
	for _, adminListPO := range adminListPOs {
		adminListDTOs = append(adminListDTOs, &service.Admin{
			ID:       adminListPO.ID,
			Name:     adminListPO.Name,
			Password: adminListPO.Password,
		})
	}
	return adminListDTOs
}

func (*UserService) CreateAdmin(ctx context.Context, req *service.CreateAdminRequest) (resp *service.CreateAdminResponse, err error) {
	resp = &service.CreateAdminResponse{}

	id, err := repository.CreateAdmin(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) UpdateAdmin(ctx context.Context, req *service.UpdateAdminRequest) (resp *service.UpdateAdminResponse, err error) {
	resp = &service.UpdateAdminResponse{}

	err = repository.UpdateAdmin(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) DeleteAdmin(ctx context.Context, req *service.DeleteAdminRequest) (resp *service.DeleteAdminResponse, err error) {
	resp = &service.DeleteAdminResponse{}

	err = repository.DeleteAdmin(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
