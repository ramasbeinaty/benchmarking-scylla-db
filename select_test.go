package main

import (
	"time"
)

// persistence date times
var start time.Time = time.Date(2023, 5, 1, 15, 00, 00, 00, time.UTC)
var end time.Time = time.Date(2023, 5, 20, 15, 00, 00, 00, time.UTC)

// func BenchmarkSelectBook(b *testing.B) {
// 	var id string
// 	var published string

// 	lgr, err := zap.NewProduction()
// 	if err != nil {
// 		panic("failed to create logger | " + err.Error())
// 	}

// 	sess, err := db.NewScyllaSession(lgr)
// 	if err != nil {
// 		panic("failed to create session | " + err.Error())
// 	}

// 	var q *gocql.Query
// 	for i := 0; i < b.N; i++ {
// 		b.ResetTimer()
// 		q = sess.Query(
// 			db.SelectBooksQuery,
// 			start,
// 			end,
// 		)
// 	}

// 	itr := q.Iter()
// 	for itr.Scan(
// 		&id,
// 		&published,
// 	) {
// 		fmt.Println("book: ", id, "|", published)
// 	}
// 	if err := itr.Close(); err != nil {
// 		panic("failed to close iter | " + err.Error())
// 	}
// }

// func BenchmarkSelectBookWithIdx(b *testing.B) {
// 	var id string
// 	var published string

// 	lgr, err := zap.NewProduction()
// 	if err != nil {
// 		panic("failed to create logger | " + err.Error())
// 	}

// 	sess, err := db.NewScyllaSession(lgr)
// 	if err != nil {
// 		panic("failed to create session | " + err.Error())
// 	}

// 	var q *gocql.Query
// 	for i := 0; i < b.N; i++ {
// 		b.ResetTimer()
// 		q = sess.Query(
// 			db.SelectBookWithIdxQuery,
// 			start,
// 			end,
// 		)
// 	}

// 	itr := q.Iter()
// 	for itr.Scan(
// 		&id,
// 		&published,
// 	) {
// 		// fmt.Println("book: ", id, "|", published)
// 	}
// 	if err := itr.Close(); err != nil {
// 		panic("failed to close iter | " + err.Error())
// 	}
// }
