package repository

import (
	"time"
	"user/constant"
	"user/internal/service"
	"user/util"

	"gorm.io/gorm"
)

type Admin struct {
	ID         string `gorm:"primarykey"`
	Name       string
	Password   string `gorm:"column:password"`
	Valid      uint
	Version    uint
	CreateTime *time.Time
	UpdateTime *time.Time
}

func CheckAdmin(req *service.LoginRequest) (bool, string, error) {
	var admin Admin
	if err := DB.Where("name=?", req.GetUsername()).Last(&admin).Error; err == gorm.ErrRecordNotFound {
		return false, "", constant.USER_NOT_EXIST
	}
	if admin.Password != req.GetPassword() {
		return false, "", constant.PASSWORD_INVALID
	}
	return true, admin.ID, nil
}

func GetAdminList(req *service.GetAdminListRequest) ([]*Admin, error) {
	var adminList = make([]*Admin, 0)
	var err error
	if req.GetSn() != "" {
		err = DB.Where("sn=?", req.GetSn()).Find(&adminList).Error
	} else if req.GetName() != "" {
		err = DB.Where("name=?", req.GetName()).Find(&adminList).Error
	} else {
		err = DB.Find(&adminList).Error
	}
	if err != nil {
		return nil, err
	}
	return adminList, nil
}

func CreateAdmin(req *service.CreateAdminRequest) (string, error) {
	var adminPO = &Admin{
		ID:         util.GetUID(),
		Name:       req.GetAdmin().GetName(),
		Password:   req.GetAdmin().GetPassword(),
		Valid:      1,
		Version:    0,
		CreateTime: util.CurrTime(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Save(adminPO).Error
	if err != nil {
		return "", err
	}
	return adminPO.ID, nil
}

func UpdateAdmin(req *service.UpdateAdminRequest) error {
	var admin = &Admin{
		ID:         req.GetAdmin().GetID(),
		Name:       req.GetAdmin().GetName(),
		Password:   req.GetAdmin().GetPassword(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Updates(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(req *service.DeleteAdminRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&Admin{}).Error
	if err != nil {
		return err
	}
	return nil
}
