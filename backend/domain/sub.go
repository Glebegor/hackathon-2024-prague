package domain

import (
	"context"
)

const (
	CollectionSub = "subs"
)

type Sub struct {
	ChannelId   int    `json:"channel_id", db:"channel_id"`
	Name        string `json:"name", db:"name"`
	Price       int    `jsom:"price", db:"price"`
	Description string `json:"description", db:"description"`
	UserId      string `json:"user_id", db:"user_id"`
	Link        string `json:"link", db:"link"`
	Images      string `json:"images", db:"images"`
	Tags        string `json:"tags", db:"tags"`
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
