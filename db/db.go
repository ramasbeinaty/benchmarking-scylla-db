package db

const KeySpace = "testScylla"

const KeySpaceCQL = "CREATE KEYSPACE IF NOT EXISTS " + KeySpace + " WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' } AND durable_writes = TRUE;"
