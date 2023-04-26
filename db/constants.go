package db

const (
	Keyspace      = "scylla"
	KeyspaceQuery = "CREATE KEYSPACE IF NOT EXISTS scylla WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' } AND durable_writes = TRUE;"
	username      = ""
	password      = ""
)

var (
	clusterIPs = []string{"127.0.0.1"}
)

const (
	InsertBookQuery = `
	INSERT INTO books (
		id,
		key,
		published_date_time,
		purchased_date_time
	) VALUES (?, ?, ?, ?)
	`
	SelectBooksQuery = `
	SELECT id,
		published_date_time
	FROM books
	WHERE purchased_date_time >= ? AND purchased_date_time < ?
`
	InsertBookWithIdxQuery = `
	INSERT INTO booksWithIdx (
		id,
		key,
		published_date_time,
		purchased_date_time
	) VALUES (?, ?, ?, ?)
	`
	SelectBooksWithIdxQuery = `
	SELECT id,
		published_date_time
	FROM booksWithIdx
	WHERE purchased_date_time >= ? AND purchased_date_time < ?
	ALLOW FILTERING
`
)
