package models

type CommentRequest struct {
	Id       string `json:"id"`
	VideoId  string `json:"video_id"`
	Text     string `json:"text"`
	UserId   string `json:"user_id"`
	CreateAt string `json:"create"`
}
