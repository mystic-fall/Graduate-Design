package dto

type GetStuListReq struct {
	Sn   string `json:"sn"`
	Name string `json:"name"`
}

type GetStuListResp struct {
	StuList []*Stu `json:"stuList"`
}

type Stu struct {
	ID       string `json:"id"`
	Sn       string `json:"sn"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}

type CreateStuReq struct {
	Student Stu `json:"student"`
}

type CreateStuResp struct {
	ID string `json:"id"`
}

type UpdateStuReq struct {
	Student Stu `json:"student"`
}

type UpdateStuResp struct {
}

type DeleteStuReq struct {
	ID string `json:"id"`
}

type DeleteStuResp struct {
}
