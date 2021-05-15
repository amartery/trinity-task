package usecase

import (
	"github.com/amartery/trinity-task/pkg/app"
	"github.com/amartery/trinity-task/pkg/app/models"
)

type AppUsecase struct {
	repository app.Repository
}

func NewAppUsecase(appRep app.Repository) app.Usecase {
	return &AppUsecase{
		repository: appRep,
	}
}

func Contains(set []string, x string) bool {
	for _, n := range set {
		if x == n {
			return true
		}
	}
	return false
}

func (use *AppUsecase) GetTodayActivityFull() ([]models.User, error) {
	// TODO валидация
	liking, err := use.repository.GetTodayActivityLikes()
	if err != nil {
		return nil, err
	}

	commenting, err := use.repository.GetTodayActivityComments()
	if err != nil {
		return nil, err
	}

	// TODO оптимизация: сравнивать по длине liking и commenting

	resultID := make([]string, 0)
	result := make([]models.User, 0)

	for _, user := range liking {
		if !Contains(resultID, user.Id) {
			resultID = append(resultID, user.Id)
			result = append(result, user)
		}
	}
	for _, user := range commenting {
		if !Contains(resultID, user.Id) {
			resultID = append(resultID, user.Id)
			result = append(result, user)
		}
	}
	return result, nil
}

func (use *AppUsecase) AddUser(u *models.User) error {
	// TODO валидация
	return use.repository.AddUser(u)
}

func (use *AppUsecase) AddComment(u *models.CommentRequest) error {
	// TODO валидация
	return use.repository.AddComment(u)
}

func (use *AppUsecase) AddLike(u *models.Lile) error {
	// TODO валидация
	return use.repository.AddLike(u)
}
