package model

type PingReq struct {
	Name string `json:"name"`
	Id   string `json:"id" binding:"required"`
}
