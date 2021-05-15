package models

type Like struct {
	VideoId  string `json:"video_id"`
	UserId   string `json:"user_id"`
	CreateAt string `json:"create"`
}
