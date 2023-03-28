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

func AdminList(ginCtx *gin.Context) {
	sn := ginCtx.Query("sn")
	name := ginCtx.Query("name")

	req := &dto.GetAdminListReq{
		Sn:   sn,
		Name: name,
	}
	rpcResp := rpc.UserRpcRepo.GetAdminList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetAdminListResp{
			AdminList: convert2AdminDTO(rpcResp.GetAdminList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2AdminDTO(AdminList []*service.Admin) []*dto.Admin {
	AdminListDTOs := make([]*dto.Admin, 0)
	for _, Admin := range AdminList {
		AdminListDTOs = append(AdminListDTOs, &dto.Admin{
			ID:       Admin.ID,
			Name:     Admin.Name,
			Password: Admin.Password,
		})
	}
	return AdminListDTOs
}

func CreateAdmin(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateAdminReq
	err = json.Unmarshal(reqBody, &req.Admin)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.CreateAdmin(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateAdminResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateAdmin(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateAdminReq
	err = json.Unmarshal(reqBody, &req.Admin)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.UpdateAdmin(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateAdminReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteAdmin(ginCtx *gin.Context) {
	var req dto.DeleteAdminReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.UserRpcRepo.DeleteAdmin(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteAdminResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
