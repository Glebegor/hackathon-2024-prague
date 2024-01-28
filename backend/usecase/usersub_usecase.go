package usecase

import (
	"context"
	"time"
	"ton-event-bot/domain"
)

type usersubUsecase struct {
	usersubRepository domain.UsersubRepository
	contextTimeout    time.Duration
}

func NewUsersubUsecase(UsersubRepository domain.UsersubRepository, timeout time.Duration) domain.UsersubUsecase {
	return &usersubUsecase{
		usersubRepository: UsersubRepository,
		contextTimeout:    timeout,
	}
}

func (tu *usersubUsecase) Create(c context.Context, usersub *domain.Usersub) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.usersubRepository.Create(ctx, usersub)
}
func (tu *usersubUsecase) Delete(c context.Context, usersubId int, userId int) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.usersubRepository.Delete(ctx, usersubId, userId)
}
func (tu *usersubUsecase) Update(c context.Context, subId int, usersubid int, sub *domain.UpdateUsersub) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.usersubRepository.Update(ctx, subId, usersubid, sub)
}
func (tu *usersubUsecase) GetById(c context.Context, subId int, usersubid int) (domain.Usersub, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.usersubRepository.GetById(ctx, subId, usersubid)
}
func (tu *usersubUsecase) GetAll(c context.Context, usersubId int) ([]domain.Usersub, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.usersubRepository.GetAll(ctx, usersubId)
}
