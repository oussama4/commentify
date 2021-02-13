package sqlite

import (
	"database/sql"

	"github.com/oussama4/commentify/config"
	_ "modernc.org/sqlite"
)

var schema string = `
	CREATE TABLE IF NOT EXISTS comments (
        id TEXT PRIMARY KEY,
        body TEXT NOT NULL,
        parent_id TEXT,
        user_id TEXT NOT NULL,
		thread_id TEXT NOT NULL,
        created_at TEXT DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (thread_id) REFERENCES threads (id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS threads (
		id TEXT PRIMARY KEY,
		url TEXT NOT NULL,
		domain TEXT NOT NULL,
		title TEXT NOT NULL
	);
`

type SqliteStore struct {
	db *sql.DB
}

func CreateStore(dbConfig config.Store) (*SqliteStore, error) {
	db, err := sql.Open("sqlite", dbConfig.Dsn)
	if err != nil {
		return nil, err
	}

	s := &SqliteStore{db}

	return s, nil
}
