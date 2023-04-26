package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func Init(sess *gocql.Session, lgr *zap.Logger) error {

	err := sess.Query(createBooksTable).Exec()
	if err != nil {
		lgr.Error("error while createBooksTable", zap.Error(err))
		return fmt.Errorf("error while createBooksTable: %w", err)
	}

	err = sess.Query(createBooksWithIdxTable).Exec()
	if err != nil {
		lgr.Error("error while createBooksWithIdxTable", zap.Error(err))
		return fmt.Errorf("error while createBooksWithIdxTable: %w", err)
	}

	err = sess.Query(createBooksIdx).Exec()
	if err != nil {
		lgr.Error("error while createBooksIdx", zap.Error(err))
		return fmt.Errorf("error while createBooksIdx: %w", err)
	}
	return nil
}

const (
	createBooksTable = `
	CREATE TABLE IF NOT EXISTS books (
		id text,
		key text,
		published_date_time timestamp,
		purchased_date_time timestamp,
		PRIMARY KEY (key, purchased_date_time, id)
		);
	`
	createBooksWithIdxTable = `
	CREATE TABLE IF NOT EXISTS booksWithIdx (
		id text,
		key text,
		published_date_time timestamp,
		purchased_date_time timestamp,
		PRIMARY KEY (key, id)
		);
`

	createBooksIdx = `
	CREATE INDEX IF NOT EXISTS on booksWithIdx (purchased_date_time);
	`
	// CREATE MATERIALIZED VIEW IF NOT EXISTS booksWithIdx_index AS
	// SELECT id text,
	// 	key text,
	// 	published_date_time text,
	// 	purchased_date_time text
	// 	FROM booksWithIdx
	// 	WHERE purchased_date_time IS NOT NULL
	// 	PRIMARY KEY (key, purchased_date_time);
	//
)

// func Init() {
// 	cluster := gocql.NewCluster("127.0.0.1") // Replace with the IP addresses of your Cassandra cluster
// 	cluster.Keyspace = "scylla"              // Connect to the system keyspace to create the new keyspace
// 	session, err := cluster.CreateSession()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Create the new keyspace
// 	err = session.Query(fmt.Sprintf(KeyspaceQuery, Keyspace)).Exec()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Connect to the new keyspace
// 	cluster.Keyspace = Keyspace
// 	session, err = cluster.CreateSession()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Perform migrations
// 	for _, migration := range migrations {
// 		err = session.Query(migration).Exec()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	fmt.Println("Keyspace and migrations created successfully!")
// }
