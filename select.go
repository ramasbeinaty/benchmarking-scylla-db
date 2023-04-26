package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/test-scylla/db"
)

func TestSelectSpeed(sess *gocql.Session) {
	start := time.Now()
	var id string
	var published string
	itr := sess.Query(
		db.SelectBooksQuery,
		"1-05-2023T15:00:00Z",
		"20-05-2023T15:00:00Z",
	).Iter()
	end := time.Since(start)

	for itr.Scan(
		&id,
		&published,
	) {
		// fmt.Println("book: ", id, "|", published)
	}
	fmt.Println("Book Execution Time - ", end)

	start1 := time.Now()
	itr1 := sess.Query(
		db.SelectBooksWithIdxQuery,
		"1-05-2023T15:00:00Z",
		"20-05-2023T15:00:00Z",
	).Iter()

	end1 := time.Since(start1)
	fmt.Println("Book With Index Execution Time - ", end1)
	for itr1.Scan(
		&id,
		&published,
	) {
		// fmt.Println("book: ", id, "|", published)
	}
	if err := itr.Close(); err != nil {
		panic("failed to close iter | " + err.Error())
	}
}
