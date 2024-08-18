package matchmaking

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/rueidis"
	"github.com/rezaAmiri123/kingscomp/steps/04_matchmaking/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/04_matchmaking/internal/repository"
	"github.com/sirupsen/logrus"
)

var (
	ErrBadRedisResponse = errors.New("bad redis response")
	ErrTimeout          = errors.New("lobby queue timeout")
)

//go:embed matchmaking.lua
var matchmakingScript string

type Matchmaking interface {
	Join(ctx context.Context, userId int64, timeout time.Duration) (entity.Lobby, bool, error)
	Leave(ctx context.Context, userId int64) error
}

var _ Matchmaking = &RedisMatchmaking{}

type RedisMatchmaking struct {
	client            rueidis.Client
	matchMakingScript *rueidis.Lua
	lobby             repository.LobbyRepository
}

func NewRedisMatchmaking(client rueidis.Client, lobby repository.LobbyRepository) *RedisMatchmaking {
	script := rueidis.NewLuaScript(matchmakingScript)
	return &RedisMatchmaking{
		client:            client,
		matchMakingScript: script,
		lobby:             lobby,
	}
}

type joinLobbyPubSubResponse struct {
	err     error
	lobbyId string
}

func (r RedisMatchmaking) Join(ctx context.Context, userId int64, timeout time.Duration) (entity.Lobby, bool, error) {
	resp, err := r.matchMakingScript.Exec(ctx, r.client,
		[]string{"matchmaking", "matchmaking"},
		[]string{"4",
			strconv.FormatInt(time.Now().Add(-time.Minute*2).Unix(), 10),
			uuid.New().String(), strconv.FormatInt(userId, 10),
			strconv.FormatInt(time.Now().Unix(), 10),
		},
	).ToArray()
	if err != nil {
		logrus.WithError(err).Errorln("couldn't join the match making")
		return entity.Lobby{}, false, err
	}

	// inside a queue, we must listen to the pub/sub
	if len(resp) == 1 {
		return entity.Lobby{}, false, nil
	}

	// you have just created a lobby
	if len(resp) == 3 {
		lobbyId, _ := resp[1].ToString()
		lobby, err := r.lobby.Get(ctx, entity.NewID("lobby", lobbyId))

		// update other users lobby id
		cmds := make([]rueidis.Completed, 0)
		for _, participant := range lobby.Participants {
			cmds = append(cmds, r.client.B().JsonSet().
				Key(entity.NewID("account", participant).String()).
				Path("$..current_lobby").Value(fmt.Sprintf(`"%s"`, lobbyId)).
				Xx().Build())
		}

		r.client.DoMulti(ctx, cmds...)

		return lobby, true, err
	}

	return entity.Lobby{}, false, ErrBadRedisResponse
}

func (r RedisMatchmaking) Leave(ctx context.Context, userId int64) error {
	//TODO implement me
	panic("implement me")
}
