package integrationtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/pkg/testhelper"
	"github.com/streadway/amqp"
)

var rabbitPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}

	pool := testhelper.StartDockerPool()

	// set up the redis container for tests
	redisRes := testhelper.StartDockerInstance(pool, "rabbitmq", "3.13.0-alpine",
		func(res *dockertest.Resource) error {
			port := res.GetPort("5672/tcp")
			con, err := amqp.Dial(fmt.Sprintf(
				"amqp://guest:guest@localhost:%s/",
				port,
			))
			if err != nil {
				return err
			}
			con.Close()
			return nil
		})
	rabbitPort = redisRes.GetPort("5672/tcp")

	// now run tests
	exitCode := m.Run()
	redisRes.Close()
	os.Exit(exitCode)
}
