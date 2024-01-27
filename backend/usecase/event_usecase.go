package repository

import (
      	"context"
	"time"
      	"ton-event-bot/domain"
)
type eventUsecase struct {
       	eventRepository domain.EventRepository 
       	contextTimeout time.Duration
}

func NewEventRepository(eventRepository domain.EventRepository, timeout time.Duration) domain.EventUsecase {
        return &eventUsecase{
		eventRepository: eventRepository,
                contextTimeout: timeout,
        }
}

func (tu *eventUsecase) Create(c context.Context, event *domain.Event) error {
        return nil
}
func (tu *eventUsecase) Delete(c context.Context) error {
        return nil
}
func (tu *eventUsecase) Update(c *gin.Context, eventId string) error {
        return nil
}
func (tu *eventUsecase) GetById(c *gin.Context, eventId string) (domain.Event) {
        var data domain.Event
        return data, nil
}
func (tu *eventUsecase) GetAll(c *gin.Context) ([]domain.Event, error){
	var data []domain.Event
        return data, nil
}

