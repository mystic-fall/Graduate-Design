package handler

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/rpc"
	"api-gateway/internal/service"
	"api-gateway/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DormitoryList(ginCtx *gin.Context) {
	sn := ginCtx.Query("sn")

	req := &dto.GetDormitoryListReq{
		Sn: sn,
	}
	rpcResp := rpc.DormitoryRpcRepo.GetDormitoryList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetDormitoryListResp{
			DormitoryList: convert2DormitoryDTO(rpcResp.GetDormitoryList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2DormitoryDTO(DormitoryList []*service.Dormitory) []*dto.Dormitory {
	DormitoryListDTOs := make([]*dto.Dormitory, 0)
	for _, Dormitory := range DormitoryList {
		DormitoryListDTOs = append(DormitoryListDTOs, &dto.Dormitory{
			ID:          Dormitory.ID,
			BuildingID:  Dormitory.BuildingID,
			Floor:       strconv.FormatUint(uint64(Dormitory.Floor), 10),
			LivedNumber: strconv.FormatUint(uint64(Dormitory.LivedNumber), 10),
			MaxNumber:   strconv.FormatUint(uint64(Dormitory.MaxNumber), 10),
			Sn:          Dormitory.Sn,
		})
	}
	return DormitoryListDTOs
}

func CreateDormitory(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateDormitoryReq
	err = json.Unmarshal(reqBody, &req.Dormitory)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.CreateDormitory(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateDormitoryResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateDormitory(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateDormitoryReq
	err = json.Unmarshal(reqBody, &req.Dormitory)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.UpdateDormitory(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateDormitoryReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteDormitory(ginCtx *gin.Context) {
	var req dto.DeleteDormitoryReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.DormitoryRpcRepo.DeleteDormitory(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteDormitoryResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
