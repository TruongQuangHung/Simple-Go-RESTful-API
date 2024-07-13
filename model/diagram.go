package model

type Diagram struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
	Code   string `json:"code"`
}
