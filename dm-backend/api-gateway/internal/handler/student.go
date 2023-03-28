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

func StudentList(ginCtx *gin.Context) {
	sn := ginCtx.Query("sn")
	name := ginCtx.Query("name")

	req := &dto.GetStuListReq{
		Sn:   sn,
		Name: name,
	}
	rpcResp := rpc.UserRpcRepo.GetStuList(req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.GetStuListResp{
			StuList: convert2StuDTO(rpcResp.GetStuList()),
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func convert2StuDTO(stuList []*service.Stu) []*dto.Stu {
	stuListDTOs := make([]*dto.Stu, 0)
	for _, stu := range stuList {
		stuListDTOs = append(stuListDTOs, &dto.Stu{
			ID:       stu.ID,
			Sn:       stu.Sn,
			Name:     stu.Name,
			Sex:      stu.Sex,
			Password: stu.Password,
		})
	}
	return stuListDTOs
}

func CreateStu(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.CreateStuReq
	err = json.Unmarshal(reqBody, &req.Student)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.CreateStu(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.CreateStuResp{
			ID: rpcResp.ID,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateStu(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UpdateStuReq
	err = json.Unmarshal(reqBody, &req.Student)
	if err != nil {
		util.LogAndPanic(err)
	}

	rpcResp := rpc.UserRpcRepo.UpdateStu(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.UpdateStuReq{},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteStu(ginCtx *gin.Context) {
	var req dto.DeleteStuReq
	id := ginCtx.Query("id")
	req.ID = id

	rpcResp := rpc.UserRpcRepo.DeleteStu(&req)

	r := util.Response{
		Code: rpcResp.GetCode(),
		Msg:  rpcResp.GetMsg(),
		Data: dto.DeleteStuResp{},
	}
	ginCtx.JSON(http.StatusOK, r)
}
