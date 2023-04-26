package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/test-scylla/db"
	"go.uber.org/zap"
)

var persistence time.Time = time.Date(2023, 5, 1, 15, 00, 00, 00, time.UTC)

func BenchmarkInsertBook(b *testing.B) {
	b.Log("benchmarking insert book")

	lgr, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger | " + err.Error())
	}

	sess, err := db.NewScyllaSession(lgr)
	if err != nil {
		panic("failed to create session | " + err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sess.Query(
			db.InsertBookQuery,
			strconv.Itoa(i),
			"key"+strconv.Itoa(i),
			"10-05-2023T15:00:00Z",
			persistence,
		).Exec()
	}

}

func BenchmarkInsertBookWithIdx(b *testing.B) {
	b.Log("benchmarking insert book with index")

	lgr, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger | " + err.Error())
	}

	sess, err := db.NewScyllaSession(lgr)
	if err != nil {
		panic("failed to create session | " + err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sess.Query(
			db.InsertBookWithIdxQuery,
			strconv.Itoa(i),
			"key"+strconv.Itoa(i),
			"10-05-2023T15:00:00Z",
			persistence,
		).Exec()
	}

}
