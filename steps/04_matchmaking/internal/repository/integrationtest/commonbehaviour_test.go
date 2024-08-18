package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/rezaAmiri123/kingscomp/steps/04_matchmaking/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/04_matchmaking/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/04_matchmaking/internal/repository/redis"
	"github.com/stretchr/testify/assert"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisclient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)

	ctx := context.Background()

	cb := repository.NewRedisCommonBehaviour[testType](redisclient)
	err = cb.Save(ctx, &testType{
		ID:   "12",
		Name: "test_12",
	})
	assert.NoError(t, err)

	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "test_13",
	})
	assert.NoError(t, err)

	val, err := cb.Get(ctx, entity.NewID("testType", "12"))
	assert.NoError(t, err)
	assert.Equal(t, "test_12", val.Name)
	assert.Equal(t, "12", val.ID)
	fmt.Println(val)

	val, err = cb.Get(ctx, entity.NewID("testType", "13"))
	assert.NoError(t, err)
	assert.Equal(t, "test_13", val.Name)
	assert.Equal(t, "13", val.ID)

	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "another_name",
	})
	assert.NoError(t, err)

	val, err = cb.Get(ctx, entity.NewID("testType", "13"))
	assert.NoError(t, err)
	assert.Equal(t, "another_name", val.Name)
	assert.Equal(t, "13", val.ID)

	val, err = cb.Get(ctx, entity.NewID("testType", "14"))
	assert.ErrorIs(t, err, repository.ErrNotFound)
}
