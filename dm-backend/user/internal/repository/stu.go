package repository

import (
	"time"
	"user/constant"
	"user/internal/service"
	"user/util"

	"gorm.io/gorm"
)

type Student struct {
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

func GetStuList(req *service.GetStuListRequest) ([]*Student, error) {
	var stuList = make([]*Student, 0)
	var err error
	if req.GetSn() != "" {
		err = DB.Where("sn=?", req.GetSn()).Find(&stuList).Error
	} else if req.GetName() != "" {
		err = DB.Where("name=?", req.GetName()).Find(&stuList).Error
	} else {
		err = DB.Find(&stuList).Error
	}
	if err != nil {
		return nil, err
	}
	return stuList, nil
}

func CreateStu(req *service.CreateStuRequest) (string, error) {
	var stuPO = &Student{
		ID:         util.GetUID(),
		Name:       req.GetStudent().GetName(),
		Password:   req.GetStudent().GetPassword(),
		Valid:      1,
		Version:    0,
		Sex:        req.GetStudent().GetSex(),
		Sn:         req.GetStudent().GetSn(),
		CreateTime: util.CurrTime(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Save(stuPO).Error
	if err != nil {
		return "", err
	}
	return stuPO.ID, nil
}

func UpdateStu(req *service.UpdateStuRequest) error {
	var stuPO = &Student{
		ID:         req.GetStudent().GetID(),
		Name:       req.GetStudent().GetName(),
		Password:   req.GetStudent().GetPassword(),
		Sex:        req.GetStudent().GetSex(),
		Sn:         req.GetStudent().GetSn(),
		UpdateTime: util.CurrTime(),
	}
	err := DB.Updates(stuPO).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteStu(req *service.DeleteStuRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&Student{}).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckStu(req *service.LoginRequest) (bool, string, error) {
	var stu Student
	if err := DB.Where("name=?", req.GetUsername()).Last(&stu).Error; err == gorm.ErrRecordNotFound {
		return false, "", constant.USER_NOT_EXIST
	}
	if stu.Password != req.GetPassword() {
		return false, "", constant.PASSWORD_INVALID
	}
	return true, stu.ID, nil
}
