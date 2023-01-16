package dto

type GetAdminListReq struct {
	Sn   string `json:"sn"`
	Name string `json:"name"`
}

type GetAdminListResp struct {
	AdminList []*Admin `json:"adminList"`
}

type Admin struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateAdminReq struct {
	Admin Admin `json:"admin"`
}

type CreateAdminResp struct {
	ID string `json:"id"`
}

type UpdateAdminReq struct {
	Admin Admin `json:"admin"`
}

type UpdateAdminResp struct {
}

type DeleteAdminReq struct {
	ID string `json:"id"`
}

type DeleteAdminResp struct {
}
