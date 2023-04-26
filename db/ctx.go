package db

import (
	"time"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func NewScyllaSession(
	lgr *zap.Logger,
) (*gocql.Session, error) {
	lgr.Info("connecting to scylla...", zap.Strings("urls", clusterIPs))
	// auth := gocql.PasswordAuthenticator{
	// 	Username: username,
	// 	Password: password,
	// }

	// cluster := CreateCluster(gocql.Quorum, "scylla", "scylla-node2", "scylla-node3")
	cluster := CreateCluster(gocql.One, "scylla", clusterIPs...)
	sess, err := gocql.NewSession(*cluster)
	if err != nil {
		lgr.Fatal("unable to connect to scylla", zap.Error(err))
		return nil, err
	}

	// - creating main session
	// cls := gocql.NewCluster(clusterIPs...)
	// cls.Authenticator = auth
	// cls.Keyspace = Keyspace
	// sess, err := cls.CreateSession()
	// if err != nil {
	// 	return nil, err
	// }

	return sess, nil
}

func CreateCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

// const (
// 	CreateKeyspace = `
//   CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {
//   	'class' : 'SimpleStrategy',
//   	'replication_factor' : '1'
//   };
//   `
// )
