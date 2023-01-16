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

func BuildingList(ginCtx *gin.Context) {
	name := ginCtx.Query("name")

	req := &dto.GetBuildingListReq{
		Name: name,
	}
	rpcResp := rpc.DormitoryRpcRepo.GetBuildingList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetBuildingListResp{
			BuildingList: convert2BuildingDTO(rpcResp.GetBuildingList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2BuildingDTO(BuildingList []*service.Building) []*dto.Building {
	BuildingListDTOs := make([]*dto.Building, 0)
	for _, Building := range BuildingList {
		BuildingListDTOs = append(BuildingListDTOs, &dto.Building{
			ID:                 Building.ID,
			DormitoryManagerID: Building.DmID,
			Location:           Building.Location,
			Name:               Building.Name,
		})
	}
	return BuildingListDTOs
}

func CreateBuilding(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateBuildingReq
	err = json.Unmarshal(reqBody, &req.Building)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.CreateBuilding(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateBuildingResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateBuilding(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateBuildingReq
	err = json.Unmarshal(reqBody, &req.Building)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.DormitoryRpcRepo.UpdateBuilding(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateBuildingReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteBuilding(ginCtx *gin.Context) {
	var req dto.DeleteBuildingReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.DormitoryRpcRepo.DeleteBuilding(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteBuildingResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
