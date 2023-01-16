package dto

type GetDMListReq struct {
	Sn   string `json:"sn"`
	Name string `json:"name"`
}

type GetDMListResp struct {
	DMList []*DM `json:"dmList"`
}

type DM struct {
	ID       string `json:"id"`
	Sn       string `json:"sn"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}

type CreateDMReq struct {
	DM DM `json:"dm"`
}

type CreateDMResp struct {
	ID string `json:"id"`
}

type UpdateDMReq struct {
	DM DM `json:"dm"`
}

type UpdateDMResp struct {
}

type DeleteDMReq struct {
	ID string `json:"id"`
}

type DeleteDMResp struct {
}
