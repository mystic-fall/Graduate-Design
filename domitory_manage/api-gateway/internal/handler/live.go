package handler

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/rpc"
	"api-gateway/internal/service"
	"api-gateway/util"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LiveList(ginCtx *gin.Context) {
	req := &dto.GetLiveListReq{}
	rpcResp := rpc.DormitoryRpcRepo.GetLiveList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetLiveListResp{
			LiveList: convert2LiveDTO(rpcResp.GetLiveList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2LiveDTO(LiveList []*service.Live) []*dto.Live {
	LiveListDTOs := make([]*dto.Live, 0)
	for _, Live := range LiveList {
		LiveListDTOs = append(LiveListDTOs, &dto.Live{
			ID:          Live.ID,
			DormitoryID: Live.DormitoryID,
			StudentID:   Live.StudentID,
			LiveDate:    Live.LiveDate,
		})
	}
	return LiveListDTOs
}

func CreateLive(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateLiveReq
	err = json.Unmarshal(reqBody, &req.Live)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.CreateLive(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateLiveResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateLive(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateLiveReq
	err = json.Unmarshal(reqBody, &req.Live)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.UpdateLive(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateLiveReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteLive(ginCtx *gin.Context) {
	var req dto.DeleteLiveReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.DormitoryRpcRepo.DeleteLive(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteLiveResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
