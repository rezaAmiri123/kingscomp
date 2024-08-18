package repository

import (
	"github.com/redis/rueidis"
	"github.com/rezaAmiri123/kingscomp/steps/05_matchmaking/internal/entity"
)
var _ LobbyRepository = &LobbyRedisRepository{}

type LobbyRedisRepository struct {
	*RedisCommonBehaviour[entity.Lobby]
}

func NewLobbyRedisRepository(client rueidis.Client) *LobbyRedisRepository {
	return &LobbyRedisRepository{
		NewRedisCommonBehaviour[entity.Lobby](client),
	}
}
