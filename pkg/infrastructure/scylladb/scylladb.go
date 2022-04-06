package scylladb

import (
	_ "base/pkg/infrastructure/environments"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig

type scyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
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

func (conn *scyllaDBConnection) createSession() {
	if session != nil {
		return
	}

	if cluster != nil {
		cluster = conn.createCluster()
	}

	sessionCreated, err := cluster.CreateSession()
	if err != nil {

	}
	session = sessionCreated
}

func NewScyllaDBConnection() *scyllaDBConnection {
	return &scyllaDBConnection{}
}

func init() {
	log.Println(os.Getenv("SCYLLA_CONSISTENCY"))
	log.Println(os.Getenv("SCYLLA_KEYSPACE"))
	log.Println(os.Getenv("SCYLLA_HOSTS"))
}
