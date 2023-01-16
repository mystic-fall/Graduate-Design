package repository

import (
	"dormitory/internal/service"
	"dormitory/util"
	"time"
)

type Live struct {
	ID          string `gorm:"primarykey"`
	DormitoryID string
	StudentID   string
	LiveDate    *time.Time
	Valid       uint
	Version     uint
	CreateTime  *time.Time
	UpdateTime  *time.Time
}

func GetLiveList(req *service.GetLiveListRequest) ([]*Live, error) {
	var LiveList = make([]*Live, 0)
	var err error
	err = DB.Find(&LiveList).Error
	if err != nil {
		return nil, err
	}
	return LiveList, nil
}

func CreateLive(req *service.CreateLiveRequest) (string, error) {
	liveDate, err := time.Parse("2006-01-02", req.GetLive().GetLiveDate())
	if err != nil {
		return "", err
	}
	var LivePO = &Live{
		ID:          util.GetUID(),
		DormitoryID: req.GetLive().GetDormitoryID(),
		StudentID:   req.GetLive().GetStudentID(),
		LiveDate:    &liveDate,
		Valid:       1,
		Version:     0,
		CreateTime:  util.CurrTime(),
		UpdateTime:  util.CurrTime(),
	}
	err = DB.Save(LivePO).Error
	if err != nil {
		return "", err
	}
	return LivePO.ID, nil
}

func UpdateLive(req *service.UpdateLiveRequest) error {
	liveDate, err := time.Parse("2006-01-02", req.GetLive().GetLiveDate())
	if err != nil {
		return err
	}
	var LivePO = &Live{
		ID:          req.GetLive().GetID(),
		DormitoryID: req.GetLive().GetDormitoryID(),
		StudentID:   req.GetLive().GetStudentID(),
		LiveDate:    &liveDate,
		UpdateTime:  util.CurrTime(),
	}
	err = DB.Updates(LivePO).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteLive(req *service.DeleteLiveRequest) error {
	err := DB.Where("id=?", req.GetID()).Delete(&Live{}).Error
	if err != nil {
		return err
	}
	return nil
}
