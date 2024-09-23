package infra

import (
	"cash/sample-cash/api/entity"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type SqlHandler interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
}

type sqlHandler struct {
	db *sqlx.DB
}

func NewInfra(db *sqlx.DB) SqlHandler {
	return &sqlHandler{db} //暗黙的　構造体を返す
}

func (s *sqlHandler) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	err := s.db.GetContext(ctx, &user, `
	SELECT
		id,
		name,
		email
		FROM
		users
		WHERE
		id = ?
		`, userID)

	if err != nil {
		log.Println(err)
		fmt.Println("error")
		fmt.Println(err)
		return nil, err
	}
	return &user, nil

}
