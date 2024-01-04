package model

type FileResponse struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Url    string `json:"url"`
	Size   int64  `json:"size"`
}
