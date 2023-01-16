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

func DMList(ginCtx *gin.Context) {
	sn := ginCtx.Query("sn")
	name := ginCtx.Query("name")

	req := &dto.GetDMListReq{
		Sn:   sn,
		Name: name,
	}
	rpcResp := rpc.UserRpcRepo.GetDMList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetDMListResp{
			DMList: convert2DMDTO(rpcResp.GetDMList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2DMDTO(DMList []*service.DM) []*dto.DM {
	DMListDTOs := make([]*dto.DM, 0)
	for _, DM := range DMList {
		DMListDTOs = append(DMListDTOs, &dto.DM{
			ID:       DM.ID,
			Sn:       DM.Sn,
			Name:     DM.Name,
			Sex:      DM.Sex,
			Password: DM.Password,
		})
	}
	return DMListDTOs
}

func CreateDM(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateDMReq
	err = json.Unmarshal(reqBody, &req.DM)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.CreateDM(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateDMResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateDM(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateDMReq
	err = json.Unmarshal(reqBody, &req.DM)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.UpdateDM(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateDMReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteDM(ginCtx *gin.Context) {
	var req dto.DeleteDMReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.UserRpcRepo.DeleteDM(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteDMResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
