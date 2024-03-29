package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/rezaAmiri123/kingscomp/internal/events"
	"github.com/rezaAmiri123/kingscomp/internal/repository/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RedisPubSubSuite struct {
	suite.Suite
	ps *events.RedisPubSub
	rc rueidis.Client

	ctx    context.Context
	cancel context.CancelFunc
	key    string
}

func TestRedisPubSub(t *testing.T) {
	rc, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	key := "lobby:1"
	s := &RedisPubSubSuite{
		ps:     events.NewRedisPubSub(ctx, rc, "lobby.*"),
		cancel: cancel,
		rc:     rc,
		ctx:    ctx,
		key:    key,
	}
	suite.Run(t, s)
}

func (r *RedisPubSubSuite) TearDownSuite() {
	r.ps.Close()
	<- time.After(time.Millisecond*100)
}

func (r *RedisPubSubSuite)TestpubSub(){
	ch := make(chan struct{})
	cancel,_ := r.ps.Register("lobby.1", events.EventAny, func(info events.EventInfo) {
		assert.Equal(r.T(),events.EventUserAnswer, info.Type)
		ch <- struct{}{}
	})
	defer cancel()

	err := r.ps.Dispatch(r.ctx,"lobby.1", events.EventUserAnswer, events.EventInfo{})
	assert.NoError(r.T(),err)
	select {
	case <-ch:
	case <-time.After(time.Second):
		r.T().Fatal("timeout message")
		
	}
}
func (r *RedisPubSubSuite)TestPubsubClose(){
	cancel, _ := r.ps.Register("lobby.1",events.EventAny, func(info events.EventInfo) {
		r.T().Fatal("this lock must not run")
	})
	cancel()

	ch := make(chan struct{})
	c2, _ := r.ps.Register("lobby.1", events.EventAny, func(info events.EventInfo) {
		ch <- struct{}{}
	})
	defer c2()
	err := r.ps.Dispatch(r.ctx, "lobby.1",events.EventUserAnswer, events.EventInfo{})
	assert.NoError(r.T(),err)
	select{
	case <-ch:
	case<-time.After(time.Millisecond*100):
		r.T().Fatal("timeout")
	}
}
