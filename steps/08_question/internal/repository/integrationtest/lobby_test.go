package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/repository/redis"
	"github.com/stretchr/testify/assert"
)

func TestLobby_Ready(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	lr := repository.NewLobbyRedisRepository(redisClient)

	err = lr.Save(ctx,entity.Lobby{
		ID:           "1",
		Participants: []int64{1, 2},
		UserState: map[int64]entity.UserState{
			1: {},
			2: {},
		},
	})
	assert.NoError(t,err)

	err = lr.UpdateUserState(ctx,"1",1,"isReady",true)
	assert.NoError(t,err)

	lobby,err := lr.Get(ctx,entity.NewID("lobby",1))
	assert.NoError(t,err)
	fmt.Println(lobby.UserState[1].IsReady)
}