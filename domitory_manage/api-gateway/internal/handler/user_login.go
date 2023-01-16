package handler

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/rpc"
	"api-gateway/util"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		util.LogAndPanic(err)
	}
	var req dto.UserLoginReq
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		util.LogAndPanic(err)
	}

	loginResp := rpc.UserRpcRepo.UserLogin(&req)
	token, err := util.GenerateToken(loginResp.UserID)
	if err != nil {
		util.LogAndPanic(err)
	}
	r := util.Response{
		Code: loginResp.GetCode(),
		Msg:  loginResp.GetMsg(),
		Data: dto.UserLoginResp{
			Token: token,
		},
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UserLogout(ginCtx *gin.Context) {
	r := util.Response{
		Code: "0000",
		Msg:  "",
		Data: nil,
	}
	ginCtx.JSON(http.StatusOK, r)
}
