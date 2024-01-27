package usecase

import (
	"context"
	"time"
	"ton-event-bot/domain"
)

type eventUsecase struct {
	eventRepository domain.EventRepository
	contextTimeout  time.Duration
}

func NewEventUsecase(eventRepository domain.EventRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUsecase{
		eventRepository: eventRepository,
		contextTimeout:  timeout,
	}
}

func (tu *eventUsecase) Create(c context.Context, event *domain.Event) error {
	return nil
}
func (tu *eventUsecase) Delete(c context.Context, eventId int, userId int) error {
	return nil
}
func (tu *eventUsecase) Update(c context.Context, eventId int, userId int, event *domain.Event) error {
	return nil
}
func (tu *eventUsecase) GetById(c context.Context, eventId int) (domain.Event, error) {
	var data domain.Event
	return data, nil
}
func (tu *eventUsecase) GetAll(c context.Context) ([]domain.Event, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.eventRepository.GetAll(ctx)
}
