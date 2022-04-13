package database

import (
	"base/pkg/application/interfaces"
	"fmt"
	"time"

	"github.com/Guilospanck/igocqlx"
	"github.com/gocql/gocql"
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

func (conn *scyllaDBConnection) createSession(cluster *gocql.ClusterConfig) (igocqlx.ISessionx, error) {
	session, err := igocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		conn.logger.Error(fmt.Sprintf("An error occurred while creating DB session: %s", err.Error()))
		return nil, err
	}

	return session, nil
}

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, logger interfaces.ILogger, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency,
		keyspace,
		logger,
		hosts,
	}
}

func GetConnection(connection *scyllaDBConnection, logger interfaces.ILogger) (igocqlx.ISessionx, error) {
	cluster := connection.createCluster()
	return connection.createSession(cluster)
}
