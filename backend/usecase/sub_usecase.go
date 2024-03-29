package usecase

import (
	"context"
	"time"
	"ton-event-bot/domain"
)

type subUsecase struct {
	subRepository  domain.SubRepository
	contextTimeout time.Duration
}

func NewSubUsecase(subRepository domain.SubRepository, timeout time.Duration) domain.SubUsecase {
	return &subUsecase{
		subRepository:  subRepository,
		contextTimeout: timeout,
	}
}

func (tu *subUsecase) Create(c context.Context, sub *domain.Sub) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.subRepository.Create(ctx, sub)
}
func (tu *subUsecase) Delete(c context.Context, subId int, userId int) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.subRepository.Delete(ctx, subId, userId)
}
func (tu *subUsecase) Update(c context.Context, subId int, sub *domain.UpdateSub) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.subRepository.Update(ctx, subId, sub)
}
func (tu *subUsecase) GetById(c context.Context, subId int) (domain.Sub, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.subRepository.GetById(ctx, subId)
}
func (tu *subUsecase) GetAll(c context.Context) ([]domain.Sub, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.subRepository.GetAll(ctx)
}
