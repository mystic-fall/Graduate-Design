package dto

type GetLiveListReq struct {
}

type GetLiveListResp struct {
	LiveList []*Live `json:"liveList"`
}

type Live struct {
	ID          string `json:"id"`
	DormitoryID string `json:"dormitoryId"`
	StudentID   string `json:"studentId"`
	LiveDate    string `json:"liveDate"`
}

type CreateLiveReq struct {
	Live Live `json:"live"`
}

type CreateLiveResp struct {
	ID string `json:"id"`
}

type UpdateLiveReq struct {
	Live Live `json:"live"`
}

type UpdateLiveResp struct {
}

type DeleteLiveReq struct {
	ID string `json:"id"`
}

type DeleteLiveResp struct {
}
