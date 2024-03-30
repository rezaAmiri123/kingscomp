package repository

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rezaAmiri123/kingscomp/internal/entity"
	"github.com/rezaAmiri123/kingscomp/pkg/jsonhelper"
)

var _ Lobby = &LobbyRedisRepository{}

type LobbyRedisRepository struct {
	*RedisCommonBehaviour[entity.Lobby]
}

func NewLobbyRedisRepository(client rueidis.Client) *LobbyRedisRepository {
	return &LobbyRedisRepository{
		NewRedisCommonBehaviour[entity.Lobby](client),
	}
}

func (l *LobbyRedisRepository) UpdateUserState(ctx context.Context, lobbyID string, UserID int64, key string, val any) error {
	updatePath := fmt.Sprintf("$.userState.%d.%s", UserID, key)
	cmd := l.client.B().JsonSet().
		Key(string(entity.NewID("lobby", lobbyID).String())).Path(updatePath).
		Value(string(jsonhelper.Encode(val))).Build()

	return l.client.Do(ctx, cmd).Error()
}
