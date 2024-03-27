package events

import (
	"context"

	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
)
type RedisQueue struct{
	rdb rueidis.Client
	
	ctx context.Context
	cancel context.CancelFunc
	key string
	inMem *InMemoryEvents
}

