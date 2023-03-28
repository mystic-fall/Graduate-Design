package repository

import (
	"time"
	"user/constant"
	"user/internal/service"
	"user/util"

	"gorm.io/gorm"
)

type DormitoryManager struct {
	ID         string `gorm:"primarykey"`
	Name       string
	Password   string `gorm:"column:password"`
	Valid      uint
	Version    uint
	Sex        string
	Sn         string
	CreateTime *time.Time
	UpdateTime *time.Time
}

func CheckDM(req *service.LoginRequest) (bool, string, error) {
	var dm DormitoryManager
	if err := DB.Where("name=?", req.GetUsername()).Last(&dm).Error; err == gorm.ErrRecordNotFound {
		return false, "", constant.USER_NOT_EXIST
	}
	if dm.Password != req.GetPassword() {
		return false, "", constant.PASSWORD_INVALID
	}
	return true, dm.ID, nil
}

func GetDMList(req *service.GetDMListRequest) ([]*DormitoryManager, error) {
	var dmList = make([]*DormitoryManager, 0)
	var err error
	if req.GetSn() != "" {
		err = DB.Where("sn=?", req.GetSn()).Find(&dmList).Error
	} else if req.GetName() != "" {
		err = DB.Where("name=?", req.GetName()).Find(&dmList).Error
	} else {
		err = DB.Find(&dmList).Error
	}
	if err != nil {
		return nil, err
	}
	return dmList, nil
}

func CreateDM(req *service.CreateDMRequest) (string, error) {
	var dmPO = &DormitoryManager{
		ID:         util.GetUID(),
		Name:       req.GetDM().GetName(),
		Password:   req.GetDM().GetPassword(),
		Valid:      1,
		Version:    0,
		Sex:        req.GetDM().GetSex(),
		Sn:         req.GetDM().GetSn(),
		CreateTime: util.CurrTime(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Save(dmPO).Error
	if err != nil {
		return "", err
	}
	return dmPO.ID, nil
}

func UpdateDM(req *service.UpdateDMRequest) error {
	var dmPO = &DormitoryManager{
		ID:         req.GetDM().GetID(),
		Name:       req.GetDM().GetName(),
		Password:   req.GetDM().GetPassword(),
		Sex:        req.GetDM().GetSex(),
		Sn:         req.GetDM().GetSn(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Updates(dmPO).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDM(req *service.DeleteDMRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&DormitoryManager{}).Error
	if err != nil {
		return err
	}
	return nil
}
