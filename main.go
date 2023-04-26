package main

import (
	"time"

	"github.com/test-scylla/db"
	"go.uber.org/zap"
)

func main() {
	lgr, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger | " + err.Error())
	}

	sess, err := db.NewScyllaSession(lgr)
	if err != nil {
		panic("failed to create session | " + err.Error())
	}

	if err := db.Init(
		sess,
		lgr,
	); err != nil {
		panic(err)
	}

	// ######### INSERT DATA INTO TABLES #############
	// keys := []int{1, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 5000}
	// keys := []int{225000, 227000}
	keys := []int{1, 3}
	persistences := []time.Time{
		time.Date(2023, 5, 8, 15, 00, 00, 00, time.UTC),
		time.Date(2023, 5, 9, 15, 00, 00, 00, time.UTC),
		time.Date(2023, 5, 10, 15, 00, 00, 00, time.UTC),
	}

	// insertBooks := `INSERT INTO books (
	// 	id,
	// 	key,
	// 	published_date_time,
	// 	purchased_date_time
	// ) VALUES ('%s', '%s', '%s', '%s');`

	// insertBooksWithIdx := `INSERT INTO booksWithIdx (
	// 	id,
	// 	key,
	// 	published_date_time,
	// 	purchased_date_time
	// ) VALUES ('%s', '%s', '%s', '%s');`

	// TestInsert(sess, keys, persistences)

	// TestManualBulkInsert(sess, keys, persistences, insertBooksWithIdx)

	TestBulkInsert(sess, keys, persistences, db.InsertBookQuery)

	// TestSelectSpeed(sess)

	// print("hi")

}
