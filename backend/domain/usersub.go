package domain

import (
	"context"
)

const (
	CollectionUsersub = "subs"
)

type Usersub struct {
	ChannelId int    `json:"channel_id" db:"channel_id"`
	UserId    string `json:"user_id" db:"user_id"`
	Endtime   string `json:"endtime" db:"endtime"`
}

type UsersubRepository interface {
	Create(c context.Context, usersub *Usersub) error
	GetById(c context.Context, usersubId int, userId int) (Usersub, error)
	GetAll(c context.Context, usersubId int) ([]Usersub, error)
	Update(c context.Context, userId int, usersubId int, usersub *UpdateUsersub) error
	Delete(c context.Context, usersubId int, userId int) error
}

type UsersubUsecase interface {
	Create(c context.Context, usersub *Usersub) error
	GetById(c context.Context, usersubId int, userId int) (Usersub, error)
	GetAll(c context.Context, usersubId int) ([]Usersub, error)
	Update(c context.Context, userId int, usersubId int, usersub *UpdateUsersub) error
	Delete(c context.Context, usersubId int, userId int) error
}
