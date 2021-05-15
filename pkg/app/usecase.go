package app

import (
	"github.com/amartery/trinity-task/pkg/app/models"
)

type Usecase interface {
	AddUser(u *models.User) error
	AddComment(u *models.CommentRequest) error
	AddLike(u *models.Like) error
	GetTodayActivityFull() ([]models.User, error)
}
