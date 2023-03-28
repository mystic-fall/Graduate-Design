package handler

import (
	"context"
	"user/constant"
	"user/internal/repository"
	"user/internal/service"
)

func (*UserService) GetStuList(ctx context.Context, req *service.GetStuListRequest) (resp *service.GetStuListResponse, err error) {
	resp = &service.GetStuListResponse{}

	stuListPO, err := repository.GetStuList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.StuList = convert2StuDTO(stuListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2StuDTO(stuListPOs []*repository.Student) []*service.Stu {
	stuListDTOs := make([]*service.Stu, 0)
	for _, stuListPO := range stuListPOs {
		stuListDTOs = append(stuListDTOs, &service.Stu{
			ID:       stuListPO.ID,
			Sn:       stuListPO.Sn,
			Name:     stuListPO.Name,
			Sex:      stuListPO.Sex,
			Password: stuListPO.Password,
		})
	}
	return stuListDTOs
}

func (*UserService) CreateStu(ctx context.Context, req *service.CreateStuRequest) (resp *service.CreateStuResponse, err error) {
	resp = &service.CreateStuResponse{}

	id, err := repository.CreateStu(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) UpdateStu(ctx context.Context, req *service.UpdateStuRequest) (resp *service.UpdateStuResponse, err error) {
	resp = &service.UpdateStuResponse{}

	err = repository.UpdateStu(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*UserService) DeleteStu(ctx context.Context, req *service.DeleteStuRequest) (resp *service.DeleteStuResponse, err error) {
	resp = &service.DeleteStuResponse{}

	err = repository.DeleteStu(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
