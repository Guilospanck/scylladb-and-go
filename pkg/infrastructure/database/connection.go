package database

import (
	"base/pkg/application/interfaces"
	_ "base/pkg/infrastructure/environments"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig

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

	clusterCreated := gocql.NewCluster(conn.hosts...)
	clusterCreated.Consistency = conn.consistency
	clusterCreated.Keyspace = conn.keyspace
	clusterCreated.Timeout = 5 * time.Second
	clusterCreated.RetryPolicy = retryPolicy
	clusterCreated.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	cluster = clusterCreated

	return clusterCreated
}

func (conn *scyllaDBConnection) createSession() (*gocql.Session, error) {
	if session != nil {
		return session, nil
	}

	if cluster != nil {
		cluster = conn.createCluster()
	}

	sessionCreated, err := cluster.CreateSession()
	if err != nil {
		conn.logger.Error("An error occurred while creating DB session: ", zap.Error(err))
		return nil, err
	}
	session = sessionCreated
	return sessionCreated, nil
}

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, logger interfaces.ILogger, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency,
		keyspace,
		logger,
		hosts,
	}
}

func GetConnection(logger interfaces.ILogger) (*gocql.Session, error) {
	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	keyspace := os.Getenv("SCYLLA_KEYSPACE")
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")

	connection := NewScyllaDBConnection(consistency, keyspace, logger, hosts...)
	return connection.createSession()
}
