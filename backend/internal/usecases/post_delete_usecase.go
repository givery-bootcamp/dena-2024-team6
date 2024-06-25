package usecases

import (
	"myapp/internal/interfaces"
)

type PostDeleteUsecase struct {
	repository interfaces.PostsRepository
}

func NewPostDeleteUsecase(r interfaces.PostsRepository) *PostDeleteUsecase {
	return &PostDeleteUsecase{
		repository: r,
	}
}

func (u *PostDeleteUsecase) Execute(userID int, postID int) error {
	return u.repository.Delete(userID, postID)
}
