package repository

import (
	"dormitory/internal/service"
	"dormitory/util"
	"time"
)

type Dormitory struct {
	ID          string `gorm:"primarykey"`
	BuildingID  string
	Floor       uint32
	LivedNumber uint32
	MaxNumber   uint32
	Sn          string
	Valid       uint
	Version     uint
	CreateTime  *time.Time
	UpdateTime  *time.Time
}

func GetDormitoryList(req *service.GetDormitoryListRequest) ([]*Dormitory, error) {
	var DormitoryList = make([]*Dormitory, 0)
	var err error
	if req.GetSn() != "" {
		err = DB.Where("sn=?", req.GetSn()).Find(&DormitoryList).Error
	} else {
		err = DB.Find(&DormitoryList).Error
	}
	if err != nil {
		return nil, err
	}
	return DormitoryList, nil
}

func CreateDormitory(req *service.CreateDormitoryRequest) (string, error) {
	var DormitoryPO = &Dormitory{
		ID:          util.GetUID(),
		BuildingID:  req.GetDormitory().GetBuildingID(),
		Floor:       req.GetDormitory().GetFloor(),
		LivedNumber: req.GetDormitory().GetLivedNumber(),
		MaxNumber:   req.GetDormitory().GetMaxNumber(),
		Sn:          req.GetDormitory().GetSn(),
		Valid:       1,
		Version:     0,
		CreateTime:  util.CurrTime(),
		UpdateTime:  util.CurrTime(),
	}
	err := DB.Save(DormitoryPO).Error
	if err != nil {
		return "", err
	}
	return DormitoryPO.ID, nil
}

func UpdateDormitory(req *service.UpdateDormitoryRequest) error {
	var DormitoryPO = &Dormitory{
		ID:          req.GetDormitory().GetID(),
		BuildingID:  req.GetDormitory().GetBuildingID(),
		Floor:       req.GetDormitory().GetFloor(),
		LivedNumber: req.GetDormitory().GetLivedNumber(),
		MaxNumber:   req.GetDormitory().GetMaxNumber(),
		Sn:          req.GetDormitory().GetSn(),
		UpdateTime:  util.CurrTime(),
	}
	err := DB.Updates(DormitoryPO).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDormitory(req *service.DeleteDormitoryRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&Dormitory{}).Error
	if err != nil {
		return err
	}
	return nil
}
