package environments

import (
	"os"

	"github.com/gocql/gocql"
)

func initializeDevEnv() {
	os.Setenv("SCYLLA_CONSISTENCY", gocql.Quorum.String())
	os.Setenv("SCYLLA_KEYSPACE", "catalog")
	os.Setenv("SCYLLA_HOSTS", "scylla-node1, scylla-node2, scylla-node3")
}

func init() {
	ENV := os.Getenv("GO_ENV")

	switch ENV {
	case "development":
		initializeDevEnv()
	default:
		initializeDevEnv()
	}
}
