package repository

import (
	"github.com/redis/rueidis"
	"github.com/rezaAmiri123/kingscomp/steps/13_pubsub/internal/entity"
)

var _ Account = &AccountRedisRepository{}

type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
