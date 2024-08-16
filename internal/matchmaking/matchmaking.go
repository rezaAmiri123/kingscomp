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
	"github.com/rezaAmiri123/kingscomp/internal/config"
	"github.com/rezaAmiri123/kingscomp/internal/entity"
	"github.com/rezaAmiri123/kingscomp/internal/repository"
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
	matchmakingScript *rueidis.Lua
	lobby             repository.Lobby
	account           repository.Account
	question          repository.Question
}

func NewRedisMatchmaking(client rueidis.Client,
	lobby repository.Lobby,
	question repository.Question,
	account repository.Account,
) *RedisMatchmaking {
	script := rueidis.NewLuaScript(matchmakingScript)
	return &RedisMatchmaking{
		client:            client,
		matchmakingScript: script,
		lobby:             lobby,
		account:           account,
		question:          question,
	}
}

func (r RedisMatchmaking) Join(ctx context.Context, userId int64, timeout time.Duration) (entity.Lobby, bool, error) {
	defer r.Leave(context.Background(), userId)

	resp, err := r.matchmakingScript.Exec(ctx, r.client,
		[]string{"matchmaking", "matchmaking"},
		[]string{fmt.Sprint(config.Default.LobbyMaxPlayer - 1),
			strconv.FormatInt(time.Now().Add(-time.Minute*2).Unix(), 10),
			uuid.New().String(), strconv.FormatInt(userId, 10),
			strconv.FormatInt(time.Now().Unix(), 10)},
	).ToArray()
	if err != nil {
		logrus.WithError(err).Errorln("couldn't join the match making")
		return entity.Lobby{}, false, err
	}

	// inside a queue, we must listen to the pub/sub
	if len(resp) == 1 {
		logrus.WithField("userId", userId).Info("waiting for a lobby")
		cmd := r.client.B().Brpop().
			Key(fmt.Sprintf("matchmaking:%d", userId)).Timeout(timeout.Seconds()).Build()
		result, err := r.client.Do(ctx, cmd).AsStrSlice()
		
	}
}

func (r RedisMatchmaking) Leave(ctx context.Context, userId int64) error {}

// func (r RedisMatchmaking){}
// func (r RedisMatchmaking){}
