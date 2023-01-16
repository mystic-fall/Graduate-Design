package repository

import (
	"dormitory/internal/service"
	"dormitory/util"
	"time"
)

type Building struct {
	ID                 string `gorm:"primarykey"`
	DormitoryManagerID string
	Location           string
	Name               string
	Valid              uint
	Version            uint
	CreateTime         *time.Time
	UpdateTime         *time.Time
}

func GetBuildingList(req *service.GetBuildingListRequest) ([]*Building, error) {
	var BuildingList = make([]*Building, 0)
	var err error
	if req.GetName() != "" {
		err = DB.Where("name=?", req.GetName()).Find(&BuildingList).Error
	} else {
		err = DB.Find(&BuildingList).Error
	}
	if err != nil {
		return nil, err
	}
	return BuildingList, nil
}

func CreateBuilding(req *service.CreateBuildingRequest) (string, error) {
	var BuildingPO = &Building{
		ID:                 util.GetUID(),
		DormitoryManagerID: req.GetBuilding().GetDmID(),
		Location:           req.GetBuilding().GetLocation(),
		Name:               req.GetBuilding().GetName(),
		Valid:              1,
		Version:            0,
		CreateTime:         util.CurrTime(),
		UpdateTime:         util.CurrTime(),
	}
	err := DB.Save(BuildingPO).Error
	if err != nil {
		return "", err
	}
	return BuildingPO.ID, nil
}

func UpdateBuilding(req *service.UpdateBuildingRequest) error {
	var BuildingPO = &Building{
		ID:                 req.GetBuilding().GetID(),
		DormitoryManagerID: req.GetBuilding().GetDmID(),
		Location:           req.GetBuilding().GetLocation(),
		Name:               req.GetBuilding().GetName(),
		UpdateTime:         util.CurrTime(),
	}
	err := DB.Updates(BuildingPO).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBuilding(req *service.DeleteBuildingRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&Building{}).Error
	if err != nil {
		return err
	}
	return nil
}
