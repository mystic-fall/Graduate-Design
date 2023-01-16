package handler

import (
	"api-gateway/constant"
	"api-gateway/internal/dto"
	"api-gateway/internal/rpc"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"

	"github.com/gin-gonic/gin"
)

func Upload(ginCtx *gin.Context) {
	file, _, err := ginCtx.Request.FormFile("file")
	if err != nil {
		log.Printf("upload file err: %v\n", err)
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	excel, err := excelize.OpenReader(file)
	if err != nil {
		log.Printf("read excel err: %v\n", err)
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 获取 Sheet1 上所有单元格
	rows, err := excel.GetRows(constant.DefaultSheet)
	if err != nil {
		fmt.Println(err)
		return
	}

	table := ginCtx.Query("table")
	HandleExcelData(table, rows)

	ginCtx.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"msg":  fmt.Sprintf("upload successful!"),
	})
}

func HandleExcelData(table string, rows [][]string) {
	switch table {
	case constant.Student:
		stuList := make([]dto.Stu, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&stuList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, stu := range stuList {
			rpc.UserRpcRepo.CreateStu(&dto.CreateStuReq{Student: stu})
		}
	case constant.DM:
		dmList := make([]dto.DM, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&dmList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, dm := range dmList {
			rpc.UserRpcRepo.CreateDM(&dto.CreateDMReq{DM: dm})
		}
	case constant.Admin:
		adminList := make([]dto.Admin, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&adminList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, admin := range adminList {
			rpc.UserRpcRepo.CreateAdmin(&dto.CreateAdminReq{Admin: admin})
		}
	case constant.Building:
		buildList := make([]dto.Building, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&buildList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, build := range buildList {
			rpc.DormitoryRpcRepo.CreateBuilding(&dto.CreateBuildingReq{Building: build})
		}
	case constant.Dormitory:
		dormitoryList := make([]dto.Dormitory, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&dormitoryList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, dormitory := range dormitoryList {
			rpc.DormitoryRpcRepo.CreateDormitory(&dto.CreateDormitoryReq{Dormitory: dormitory})
		}
	case constant.Live:
		liveList := make([]dto.Live, len(rows))
		for i := 1; i < len(rows); i++ {
			for j := 0; j < len(rows[0]); j++ {
				reflect.ValueOf(&liveList[i]).Elem().Field(j + 1).SetString(rows[i][j])
			}
		}
		for _, live := range liveList {
			rpc.DormitoryRpcRepo.CreateLive(&dto.CreateLiveReq{Live: live})
		}
	}
}

func Download(ginCtx *gin.Context) {
	excel := excelize.NewFile()
	sheet := excel.NewSheet(constant.DefaultSheet)
	// 设置默认打开的表单
	excel.SetActiveSheet(sheet)

	table := ginCtx.Query("table")
	excelHeaderMap := constant.HeaderMaps[table]
	for axis, field := range excelHeaderMap {
		excel.SetCellValue(constant.DefaultSheet, axis, field)
	}
	switch table {
	case constant.Student:
		req := &dto.GetStuListReq{}
		rpcResp := rpc.UserRpcRepo.GetStuList(req)
		stuList := convert2StuDTO(rpcResp.GetStuList())
		WriteInExcel(excel, stuList)
	case constant.DM:
		req := &dto.GetDMListReq{}
		rpcResp := rpc.UserRpcRepo.GetDMList(req)
		stuList := convert2DMDTO(rpcResp.GetDMList())
		WriteInExcel(excel, stuList)
	case constant.Admin:
		req := &dto.GetAdminListReq{}
		rpcResp := rpc.UserRpcRepo.GetAdminList(req)
		stuList := convert2AdminDTO(rpcResp.GetAdminList())
		WriteInExcel(excel, stuList)
	case constant.Building:
		req := &dto.GetBuildingListReq{}
		rpcResp := rpc.DormitoryRpcRepo.GetBuildingList(req)
		stuList := convert2BuildingDTO(rpcResp.GetBuildingList())
		WriteInExcel(excel, stuList)
	case constant.Dormitory:
		req := &dto.GetDormitoryListReq{}
		rpcResp := rpc.DormitoryRpcRepo.GetDormitoryList(req)
		stuList := convert2DormitoryDTO(rpcResp.GetDormitoryList())
		WriteInExcel(excel, stuList)
	case constant.Live:
		req := &dto.GetLiveListReq{}
		rpcResp := rpc.DormitoryRpcRepo.GetLiveList(req)
		stuList := convert2LiveDTO(rpcResp.GetLiveList())
		WriteInExcel(excel, stuList)
	}

	buf, err := excel.WriteToBuffer()
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"message": "read failed",
		})
	}

	ginCtx.Set("content-type", "application/msexcel")
	ginCtx.Header("Content-disposition", "attachment; filename=download.xls")
	ginCtx.Data(http.StatusOK, "application/msexcel", buf.Bytes())
}

func WriteInExcel(excel *excelize.File, dataList interface{}) {
	var axisX rune = 'A'
	var axisY int = 2
	var axis string
	val := reflect.ValueOf(dataList)
	for i := 0; i < val.Len(); i++ {
		data := val.Index(i)
		for j := 0; j < data.Elem().NumField(); j++ {
			field := data.Elem().Field(j).String()
			axis = strconv.QuoteRuneToASCII(axisX + rune(j))[1:2] + strconv.FormatInt(int64(axisY+i), 10)
			excel.SetCellValue(constant.DefaultSheet, axis, field)
		}
	}
	return
}
