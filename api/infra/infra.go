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
	db    *sqlx.DB
	cache *cache.Cache
}

func NewInfra(db *sqlx.DB) SqlHandler {
	// キャッシュの初期化をここで行い、構造体にセット
	c := cache.New(30*time.Second, 10*time.Minute)
	return &sqlHandler{db, c}
}

func (s *sqlHandler) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User

	key := fmt.Sprintf("id_%s", userID)

	// キャッシュから取得
	if x, found := s.cache.Get(key); found {
		fmt.Println("got from cache")
		return x.(*entity.User), nil
	}

	// データベースから取得
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

	// キャッシュにセット
	s.cache.Set(key, &user, cache.DefaultExpiration)
	fmt.Println("got from db")

	return &user, nil
}
