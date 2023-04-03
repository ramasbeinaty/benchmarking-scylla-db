package db

import (
	"fmt"

	"github.com/gocql/gocql"
)

const (
	keyspace      = "test"
	keyspaceQuery = "CREATE KEYSPACE IF NOT EXISTS " + KeySpace + " WITH replication = { 'class': 'SimpleStrategy', 'replication_factor': '1' } AND durable_writes = TRUE;"
)

func Init() {
	cluster := gocql.NewCluster("127.0.0.1") // Replace with the IP addresses of your Cassandra cluster
	cluster.Keyspace = "system"              // Connect to the system keyspace to create the new keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Create the new keyspace
	err = session.Query(fmt.Sprintf(keyspaceQuery, keyspace)).Exec()
	if err != nil {
		panic(err)
	}

	// Connect to the new keyspace
	cluster.Keyspace = keyspace
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Perform migrations
	for _, migration := range migrations {
		err = session.Query(migration).Exec()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Keyspace and migrations created successfully!")
}
