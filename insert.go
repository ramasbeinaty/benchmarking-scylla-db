package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	"github.com/test-scylla/db"
)

func TestInsert(sess *gocql.Session, keys []int, persistences []string) {
	rows := 0
	counter := 0
	for _, key := range keys {
		i := rand.Intn(3)
		counter = rows

		// for idx := 1; counter < key; idx++ {
		// 	if err := sess.Query(
		// 		db.InsertBookQuery,
		// 		strconv.Itoa(counter),
		// 		"key"+strconv.Itoa(counter),
		// 		strconv.Itoa(idx)+"-05-2023T15:00:00Z",
		// 		persistences[i],
		// 	).Exec(); err != nil {
		// 		panic("failed to insert user into scylla | " + err.Error())
		// 	}
		// 	counter++
		// }

		counter = rows
		for idx := 1; counter < key; idx++ {
			if err := sess.Query(
				db.InsertBookWithIdxQuery,
				strconv.Itoa(counter),
				"key"+strconv.Itoa(counter),
				strconv.Itoa(idx)+"-05-2023T15:00:00Z",
				persistences[i],
			).Exec(); err != nil {
				panic("failed to insert user into scylla | " + err.Error())
			}
			counter++
		}

		rows = key
	}
}

func TestBulkInsert(sess *gocql.Session, keys []int, persistences []time.Time, insertStmnt string) {

	publishedDate := time.Date(2023, 5, 5, 15, 00, 00, 00, time.UTC)

	idx := 0
	for key := keys[0]; key < keys[1]; key++ {
		// in the same key/partition
		// start batch
		batch := sess.NewBatch(gocql.UnloggedBatch)

		for i := 0; i < len(persistences); i++ {

			batch.Query(
				insertStmnt,
				strconv.Itoa(idx),
				strconv.Itoa(key),
				publishedDate,
				persistences[i],
			)

			idx++
		}

		// execute batch
		if err := sess.ExecuteBatch(batch); err != nil {
			panic("failed to bulk insert -- " + err.Error())
		}

	}

}

func TestManualBulkInsert(sess *gocql.Session, keys []int, persistences []time.Time, insertStmnt string) {
	// Key property is the partition

	publishedDate := time.Date(2023, 5, 5, 15, 00, 00, 00, time.UTC)

	// partitionedBatchStatements := map[string][]string{}
	//// "key ": ""
	partitionedBatchStatements := []string{}

	idx := 0
	for key := keys[0]; key < keys[1]; key++ {
		// in the same key/partition
		stmnt := "BEGIN BATCH\n"

		for i := 0; i < len(persistences); i++ {
			stmnt += fmt.Sprintf(
				insertStmnt+"\n",
				strconv.Itoa(idx),
				strconv.Itoa(key),
				publishedDate.Format("2006-01-02 15:04:05"),
				persistences[i].Format("2006-01-02 15:04:05"),
			)

			fmt.Println("stmnt: ", stmnt)

			idx++
		}

		// end batch
		stmnt += "\nAPPLY BATCH;"

		if err := sess.Query(stmnt).Exec(); err != nil {
			panic("failed to bulk insert -- " + err.Error())
		}

		partitionedBatchStatements = append(partitionedBatchStatements, stmnt)
	}

	fmt.Println("partitioned batch statements\n", partitionedBatchStatements[0])
	fmt.Println(len(partitionedBatchStatements))

}
