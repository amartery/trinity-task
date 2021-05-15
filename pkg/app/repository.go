package app

import (
	"github.com/amartery/trinity-task/pkg/app/models"
)

type Repository interface {
	AddUser(u *models.User) error
	AddComment(u *models.CommentRequest) error
	AddLike(u *models.Lile) error
	GetTodayActivityComments() ([]models.User, error)
	GetTodayActivityLikes() ([]models.User, error)
}
