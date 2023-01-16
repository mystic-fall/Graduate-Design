package dto

type GetDormitoryListReq struct {
	Sn string `json:"sn"`
}

type GetDormitoryListResp struct {
	DormitoryList []*Dormitory `json:"dormitoryList"`
}

type Dormitory struct {
	ID          string `json:"id"`
	BuildingID  string `json:"buildingId"`
	Floor       string `json:"floor"`
	LivedNumber string `json:"livedNumber"`
	MaxNumber   string `json:"maxNumber"`
	Sn          string `json:"sn"`
}

type CreateDormitoryReq struct {
	Dormitory Dormitory `json:"dormitory"`
}

type CreateDormitoryResp struct {
	ID string `json:"id"`
}

type UpdateDormitoryReq struct {
	Dormitory Dormitory `json:"dormitory"`
}

type UpdateDormitoryResp struct {
}

type DeleteDormitoryReq struct {
	ID string `json:"id"`
}

type DeleteDormitoryResp struct {
}
