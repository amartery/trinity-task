package models

type CommentRequest struct {
	Id       string `json:"id"`
	VideoId  string `json:"video_id"`
	Text     string `json:"text"`
	UserId   string `json:"user_id"`
	CreateAt string `json:"create"`
}

var (
	CommentForSwagger = CommentRequest{VideoId: "12", Text: "very cool", UserId: "2", CreateAt: "2021-05-15 12:00:00"}
)
