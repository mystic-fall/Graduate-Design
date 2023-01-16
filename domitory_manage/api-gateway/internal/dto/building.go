package dto

type GetBuildingListReq struct {
	Name string `json:"name"`
}

type GetBuildingListResp struct {
	BuildingList []*Building `json:"buildingList"`
}

type Building struct {
	ID                 string `json:"id"`
	DormitoryManagerID string `json:"dormitoryManagerId"`
	Location           string `json:"location"`
	Name               string `json:"name"`
}

type CreateBuildingReq struct {
	Building Building `json:"building"`
}

type CreateBuildingResp struct {
	ID string `json:"id"`
}

type UpdateBuildingReq struct {
	Building Building `json:"building"`
}

type UpdateBuildingResp struct {
}

type DeleteBuildingReq struct {
	ID string `json:"id"`
}

type DeleteBuildingResp struct {
}
