package handler

import (
	"context"
	"dormitory/constant"
	"dormitory/internal/repository"
	"dormitory/internal/service"
)

func (*DormitoryService) GetLiveList(ctx context.Context, req *service.GetLiveListRequest) (resp *service.GetLiveListResponse, err error) {
	resp = &service.GetLiveListResponse{}

	LiveListPO, err := repository.GetLiveList(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.LiveList = convert2LiveDTO(LiveListPO)
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func convert2LiveDTO(LiveListPOs []*repository.Live) []*service.Live {
	LiveListDTOs := make([]*service.Live, 0)
	for _, LiveListPO := range LiveListPOs {
		LiveListDTOs = append(LiveListDTOs, &service.Live{
			ID:          LiveListPO.ID,
			DormitoryID: LiveListPO.DormitoryID,
			StudentID:   LiveListPO.StudentID,
			LiveDate:    LiveListPO.LiveDate.Format("2006-01-02"),
		})
	}
	return LiveListDTOs
}

func (*DormitoryService) CreateLive(ctx context.Context, req *service.CreateLiveRequest) (resp *service.CreateLiveResponse, err error) {
	resp = &service.CreateLiveResponse{}

	id, err := repository.CreateLive(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.ID = id
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) UpdateLive(ctx context.Context, req *service.UpdateLiveRequest) (resp *service.UpdateLiveResponse, err error) {
	resp = &service.UpdateLiveResponse{}

	err = repository.UpdateLive(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}

func (*DormitoryService) DeleteLive(ctx context.Context, req *service.DeleteLiveRequest) (resp *service.DeleteLiveResponse, err error) {
	resp = &service.DeleteLiveResponse{}

	err = repository.DeleteLive(req)
	if err != nil {
		resp.Code = constant.INNER_SERVER_ERR.Code
		resp.Msg = constant.INNER_SERVER_ERR.Msg
	}
	resp.Code = constant.SUCCESS_CODE
	resp.Msg = constant.SUCCESS_MSG
	return resp, nil
}
