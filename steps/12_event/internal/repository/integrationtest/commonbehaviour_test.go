package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/repository/redis"
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
	err = cb.Save(ctx, testType{
		ID:   "12",
		Name: "test_12",
	})
	assert.NoError(t, err)

	err = cb.Save(ctx, testType{
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

	err = cb.Save(ctx, testType{
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

func TestCommonBehaviourMGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	for i := 0; i < 10; i++ {
		err = cb.Save(ctx, testType{
			ID:   strconv.Itoa(i),
			Name: "Behzad",
		})
		assert.NoError(t, err)
	}

	items, err := cb.MGet(ctx,
		entity.NewID("testType", 2),
		entity.NewID("testType", 3),
		entity.NewID("testType", 4),
	)
	assert.NoError(t, err)
	assert.Len(t, items, 3)
	assert.Equal(t, "Behzad", items[0].Name)
}

func TestCommonBehaviourMGetNotExists(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	items, err := cb.MGet(ctx,
		entity.NewID("testType", -100),
		entity.NewID("testType", -200),
		entity.NewID("testType", -300),
	)

	assert.NoError(t, err)
	assert.Len(t, items, 0)

}

func TestCommonBehaviourMSet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	err = cb.MSet(ctx,
		testType{ID: "er1", Name: "Name"},
		testType{ID: "er2", Name: "Name"},
		testType{ID: "er3", Name: "Name"},
	)
	assert.NoError(t, err)

	gets, err := cb.MGet(ctx,
		entity.NewID("testType", "er1"),
		entity.NewID("testType", "er2"),
		entity.NewID("testType", "er3"),
	)
	assert.NoError(t, err)
	assert.Len(t, gets, 3)
	assert.Equal(t, gets[0].ID, "er1")
	assert.Equal(t, gets[0].Name, "Name")
}

func TestCommonBehaviourUploadField(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	err = cb.MSet(ctx,
		testType{ID: "a1", Name: "Name"},
	)
	assert.NoError(t, err)

	err = cb.SetField(ctx, entity.NewID("testType", "a1"), "Name", "Test")
	assert.NoError(t, err)

	ent, err := cb.Get(ctx, entity.NewID("testType", "a1"))
	assert.NoError(t, err)
	assert.Equal(t, ent.Name, "Test")
}

func TestCommonBehaviourKeys(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	err = cb.MSet(ctx,
		testType{ID: "a1", Name: "Amir"},
	)

	keys, err := cb.AllIDs(ctx, "testType")
	assert.NoError(t, err)
	assert.Contains(t, keys[0], "testType:")
}
