package integrationtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/repository/redis"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/pkg/testhelper"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}

	pool := testhelper.StartDockerPool()

	// set up the redis container for tests
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		})

	redisPort = redisRes.GetPort("6379/tcp")

	// now run tests
	exitCode := m.Run()
	redisRes.Close()
	os.Exit(exitCode)
}
