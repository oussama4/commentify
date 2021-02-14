package sqlite

import (
	"database/sql"

	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/model"
	"github.com/oussama4/commentify/store"
	_ "modernc.org/sqlite"
)

var schema string = `
	CREATE TABLE IF NOT EXISTS comments (
        Id TEXT PRIMARY KEY,
        Body TEXT NOT NULL,
        ParentId TEXT,
        UserId TEXT NOT NULL,
		ThreadId TEXT NOT NULL,
        CreatedAt TEXT DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (UserId) REFERENCES users (Id) ON DELETE CASCADE,
        FOREIGN KEY (ThreadId) REFERENCES threads (Id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS users (
		Id TEXT PRIMARY KEY,
		Name TEXT NOT NULL,
		Email TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS threads (
		Id TEXT PRIMARY KEY,
		Url TEXT NOT NULL,
		Domain TEXT NOT NULL,
		Title TEXT NOT NULL
	);
`

type SqliteStore struct {
	db *sql.DB
}

func CreateStore(dbConfig config.Store) (store.Store, error) {
	db, err := sql.Open("sqlite", dbConfig.Dsn)
	if err != nil {
		return nil, err
	}

	s := &SqliteStore{db}

	return s, nil
}

func (s *SqliteStore) GetComment(id string) (*model.Comment, error) {
	q := "SELECT * FROM comments WHERE Id=?"
	c := &model.Comment{}
	err := s.db.QueryRow(q, id).Scan(&c.Id, &c.Body, &c.ParentId, &c.UserId, &c.ThreadId, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.NewErrNotFound("comment")
		}
		return nil, err
	}
	return c, nil
}

func (s *SqliteStore) CreateThread(url, domain, title string) (string, error) {
	q := "INSERT INTO threads(Id, Url, Domain, Title) VALUES(?, ?, ?, ?)"
	id := store.Uid()
	_, err := s.db.Exec(q, &id, &url, &domain, &title)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *SqliteStore) CreateUser(name, email string) (string, error) {
	q := "INSERT INTO users(Id, Name, Email) VALUES(?, ?, ?)"
	id := store.Uid()
	_, err := s.db.Exec(q, &id, &name, &email)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *SqliteStore) CreateComment(body, parentId, userId, threadId string) (string, error) {
	q := "INSERT INTO comments(Id, Body, ParentId, UserId, ThreadId) VALUES(?, ?, ?, ?, ?)"
	id := store.Uid()
	_, err := s.db.Exec(q, &id, &body, &parentId, &userId, &threadId)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *SqliteStore) UpdateComment(id, body string) error {
	q := "UPDATE comments WHERE Id=? set Body=?"
	_, err := s.db.Exec(q, &id, &body)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteComment(id string) error {
	q := "DELETE FROM comments WHERE Id=?"
	_, err := s.db.Exec(q, &id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) GetThread(id string) (*model.Thread, error) {
	q := "SELECT * FROM threads WHERE Id=?"
	t := &model.Thread{}
	err := s.db.QueryRow(q, &id).Scan(&t.Id, &t.Url, &t.Domain, &t.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.NewErrNotFound("thread")
		}
		return nil, err
	}

	return t, nil
}

func (s *SqliteStore) ListComments(threadId string) ([]*model.CommentOutput, error) {
	q := `SELECT c.Id, c.Body, c.CreatedAt, u.Id, u.Name, u.Email
		FROM comments c
		INNER JOIN users u ON c.UserId = u.Id
		WHERE c.threadId=?`

	return s.listComments(q, &threadId)
}

func (s *SqliteStore) ListChildComments(threadId, parentId string) ([]*model.CommentOutput, error) {
	q := `SELECT c.Id, c.Body, c.CreatedAt, u.Id, u.Name, u.Email
		FROM comments c
		INNER JOIN users u ON c.UserId = u.Id
		WHERE c.ThreadId=? AND c.ParentId=?`

	return s.listComments(q, &threadId, &parentId)
}

func (s *SqliteStore) listComments(query string, dest ...interface{}) ([]*model.CommentOutput, error) {
	rows, err := s.db.Query(query, dest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]*model.CommentOutput, 1)
	for rows.Next() {
		o := &model.CommentOutput{}
		if err := rows.Scan(&o.Id, &o.Body, &o.CreatedAt, &o.Author.Id, &o.Author.Name, &o.Author.Email); err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return out, nil
}
