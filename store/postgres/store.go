package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/oussama4/commentify/model"
	"github.com/oussama4/commentify/store"
	sb "github.com/oussama4/sqlbuilder"
	"github.com/oussama4/stx/crypto"
)

type PostgresStore struct {
	db *sql.DB
}

func Create(dsn string) (store.Store, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	s := &PostgresStore{db: db}
	return s, nil
}

func (s *PostgresStore) GetComment(id string) (*model.Comment, error) {
	q, args := sb.Select().From("comments").Where(sb.Eq("id", id)).Query()
	c := &model.Comment{}
	err := s.db.QueryRow(q, args...).Scan(&c.Id, &c.Body, &c.ParentId, &c.UserId, &c.ThreadId, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return c, nil
}

func (s *PostgresStore) CreateThread(url, domain, title string) (string, error) {
	id := crypto.Token(21)
	q, args := sb.Insert("threads").
		Columns("id", "url", "domain", "title").
		Values(&id, &url, &domain, &title).
		Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PostgresStore) ListUsers(page int, pageSize int) ([]model.User, error) {
	q, args := sb.Select().From("users").Limit(pageSize).Offset((page - 1) * pageSize).Query()
	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.User, 0)
	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Admin); err != nil {
			return nil, err
		}
		out = append(out, u)
	}

	return out, nil
}

func (s *PostgresStore) GetUser(id string) (*model.User, error) {
	q, args := sb.Select().From("users").Where(sb.Eq("id", id)).Query()
	u := &model.User{}
	err := s.db.QueryRow(q, args...).Scan(&u.Id, &u.Name, &u.Email, &u.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return u, nil
}

func (s *PostgresStore) CreateUser(name, email string) (string, error) {
	id := crypto.Token(21)
	q, args := sb.Insert("users").Columns("id", "name", "email").Values(&id, &name, &email).Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PostgresStore) CreateComment(body, parentId, userId, threadId string) (string, error) {
	id := crypto.Token(21)
	q, args := sb.Insert("comments").
		Columns("id", "body", "parent_id", "user_id", "thread_id").
		Values(&id, &body, &parentId, &userId, &threadId).
		Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PostgresStore) UpdateComment(id, body string) error {
	q, args := sb.Update("comments").Where(sb.Eq("id", &id)).Set("body", &body).Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteComment(id string) error {
	q, args := sb.DeleteFrom("comments").Where(sb.Eq("id", &id)).Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetThread(id string) (*model.Thread, error) {
	q, args := sb.Select().From("threads").Where(sb.Eq("id", id)).Query()
	t := &model.Thread{}
	err := s.db.QueryRow(q, args...).Scan(&t.Id, &t.Url, &t.Domain, &t.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *PostgresStore) ListThreads(page int, pageSize int) ([]model.Thread, error) {
	q, args := sb.Select("id", "url", "domain", "title").
		From("threads").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Query()
	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.Thread, 0)
	for rows.Next() {
		t := model.Thread{}
		if err := rows.Scan(&t.Id, &t.Url, &t.Domain, &t.Title); err != nil {
			return nil, err
		}
		out = append(out, t)
	}

	return out, nil
}

func (s *PostgresStore) ListComments(threadId string, parentId string, page int, pageSize int) ([]model.Comment, error) {
	qb := sb.Select().From("comments")

	if threadId != "" && parentId == "" {
		q, args := qb.Where(sb.Eq("thread_id", threadId)).Limit(pageSize).Offset((page - 1) * pageSize).Query()
		return s.listComments(q, args...)
	} else if parentId != "" {
		q, args := qb.Where(sb.And(
			sb.Eq("thread_id", threadId),
			sb.Eq("parent_id", parentId),
		)).Limit(pageSize).Offset((page - 1) * pageSize).Query()
		return s.listComments(q, args...)
	}
	q, args := qb.Limit(pageSize).Offset((page - 1) * pageSize).Query()
	return s.listComments(q, args...)
}

func (s *PostgresStore) listComments(query string, dest ...interface{}) ([]model.Comment, error) {
	rows, err := s.db.Query(query, dest...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.Comment, 0)
	for rows.Next() {
		o := model.Comment{}
		if err := rows.Scan(&o.Id, &o.Body, &o.ParentId, &o.UserId, &o.ThreadId, &o.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return out, nil
}

func (s *PostgresStore) CountComments(threadId string) (int, error) {
	qb := sb.Select("COUNT(*)").From("comments")
	c := -1
	if threadId != "" {
		q, args := qb.Where(sb.Eq("thread_id", threadId)).Query()
		if err := s.db.QueryRow(q, args...).Scan(&c); err != nil {
			return c, err
		}
		return c, nil
	}
	q, _ := qb.Query()
	if err := s.db.QueryRow(q).Scan(&c); err != nil {
		return c, err
	}
	return c, nil
}
