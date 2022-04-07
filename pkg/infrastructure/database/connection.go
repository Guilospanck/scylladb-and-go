package database

import (
	"base/pkg/application/interfaces"
	_ "base/pkg/infrastructure/environments"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"go.uber.org/zap"
)

type scyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	logger      interfaces.ILogger
	hosts       []string
}

func (conn *scyllaDBConnection) createCluster() *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}

	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	return cluster
}

func (conn *scyllaDBConnection) createSession(cluster *gocql.ClusterConfig) (*gocqlx.Session, error) {
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		conn.logger.Error("An error occurred while creating DB session: ", zap.Error(err))
		return nil, err
	}

	return &session, nil
}

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, logger interfaces.ILogger, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency,
		keyspace,
		logger,
		hosts,
	}
}

func GetConnection(logger interfaces.ILogger) (*gocqlx.Session, error) {
	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	keyspace := os.Getenv("SCYLLA_KEYSPACE")
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")

	connection := NewScyllaDBConnection(consistency, keyspace, logger, hosts...)
	cluster := connection.createCluster()
	return connection.createSession(cluster)
}
