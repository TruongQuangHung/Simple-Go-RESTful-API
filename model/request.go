package model

type AddUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AddDiagramRequest struct {
	UserId int    `json:"userId"`
	Type   string `json:"type"`
	Code   string `json:"code"`
}

type UpdateDiagramRequest struct {
	UserId *int    `json:"userId,omitempty"`
	Type   *string `json:"type,omitempty"`
	Code   *string `json:"code,omitempty"`
}
