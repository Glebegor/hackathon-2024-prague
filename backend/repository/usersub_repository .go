package repository

import (
	"context"
	"fmt"

	"ton-event-bot/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type usersubRepository struct {
	database   *sqlx.DB
	collection string
	Table      string
}

func NewUsersubRepository(db *sqlx.DB, collection string) domain.UsersubRepository {
	return &usersubRepository{
		database:   db,
		collection: collection,
		Table:      "sub_users",
	}
}

func (tr *usersubRepository) Create(c context.Context, usersub *domain.Usersub) error {
	query := fmt.Sprintf("INSERT INTO %s (channel_id, user_id, endtime) VALUES ($1, $2, $3)", tr.Table)
	_, err := tr.database.Exec(query, usersub.ChannelId, usersub.UserId, usersub.Endtime)
	if err != nil {
		return err
	}
	return nil
}
func (tr *usersubRepository) Delete(c context.Context, usersubId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE channel_id=$1 AND user_id=$2", tr.Table)
	_, err := tr.database.Exec(query, usersubId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (tr *usersubRepository) Update(c context.Context, usersubId int, userId int, usersub *domain.UpdateUsersub) error {
	query := fmt.Sprintf("UPDATE %s SET endtime=%s WHERE channel_id=%d and user_id=$1", tr.Table, usersub.Endtime, usersubId)
	fmt.Print(query, userId)
	_, err := tr.database.Exec(query)
	return err
}

func (tr *usersubRepository) GetById(c context.Context, usersubId int, userId int) (domain.Usersub, error) {
	var data domain.Usersub
	query := fmt.Sprintf("SELECT * FROM %s WHERE channel_id=$1 and user_id=$2", tr.Table)
	err := tr.database.QueryRow(query, usersubId, userId).Scan(&data.ChannelId, &data.UserId, &data.Endtime)

	return data, err
}

func (tr *usersubRepository) GetAll(c context.Context, usersubId int) ([]domain.Usersub, error) {
	var data []domain.Usersub

	query := fmt.Sprintf("SELECT * from %s WHERE user_id=$1", tr.Table)
	rows, err := tr.database.Query(query, usersubId)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var dataEl domain.Usersub
		err := rows.Scan(&dataEl.ChannelId, &dataEl.UserId, &dataEl.Endtime)
		if err != nil {
			return data, err
		}
		data = append(data, dataEl)
	}
	return data, nil
}
