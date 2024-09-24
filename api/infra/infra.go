package infra

import (
	"cash/sample-cash/api/entity"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
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

	// キャッシュの初期設定
	c := cache.New(30*time.Second, 10*time.Minute)

	key := fmt.Sprintf("id_%s", userID)

	if x, found := c.Get(key); found {
		fmt.Println("got from cash")
		return x.(*entity.User), nil
	}

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

	c.Set(key, &user, cache.DefaultExpiration)
	fmt.Println("got from db")

	return &user, nil

}
