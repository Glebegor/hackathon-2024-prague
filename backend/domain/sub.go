package domain

import (
	"context"
)

const (
	CollectionSub = "subs"
)

type Sub struct {
	ChannelId   string `json:"channel_id"`
	Name        string `json:"name"`
	Price       string `jsom:"price"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	Images      string `json:"images"`
	Tags        string `json:"tags"`
}

type SubRepository interface {
	Create(c context.Context, sub *Sub) error
	GetById(c context.Context, subId int) (Sub, error)
	GetAll(c context.Context) ([]Sub, error)
	Update(c context.Context, subId int, userId int, sub *Sub) error
	Delete(c context.Context, subId int, userId int) error
}

type SubUsecase interface {
	Create(c context.Context, sub *Sub) error
	GetById(c context.Context, subId int) (Sub, error)
	GetAll(c context.Context) ([]Sub, error)
	Update(c context.Context, subId int, userId int, sub *Sub) error
	Delete(c context.Context, subId int, userId int) error
}
