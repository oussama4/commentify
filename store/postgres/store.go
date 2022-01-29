package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/oussama4/commentify/model"
	"github.com/oussama4/commentify/store"
	sb "github.com/oussama4/sqlbuilder"
	"github.com/oussama4/stx/crypto"
)

type PostgresStore struct {
	db *sql.DB
}

func Create(dsn string) (store.Store, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	s := &PostgresStore{db: db}
	return s, nil
}

func (s *PostgresStore) GetComment(id string) (*model.Comment, error) {
	q, args := sb.Select().From("comments").Where(sb.Eq("id", id)).Query()
	c := &model.Comment{}
	err := s.db.QueryRow(q, args...).Scan(&c.Id, &c.Body, &c.ParentId, &c.UserId, &c.PageId, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return c, nil
}

func (s *PostgresStore) CreatePage(url, title string) (string, error) {
	id := crypto.Token(21)
	q, args := sb.Insert("pages").
		Columns("id", "url", "title").
		Values(&id, &url, &title).
		Query()
	_, err := s.db.Exec(q, args...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PostgresStore) ListUsers(page int, pageSize int) ([]model.User, error) {
	q, args := sb.Select("id", "name", "email").From("users").
		Limit(pageSize).Offset((page - 1) * pageSize).Query()
	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.User, 0)
	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		out = append(out, u)
	}

	return out, nil
}

func (s *PostgresStore) GetUser(id string) (*model.User, error) {
	q, args := sb.Select().From("users").Where(sb.Eq("id", id)).Query()
	u := &model.User{}
	err := s.db.QueryRow(q, args...).Scan(&u.Id, &u.Name, &u.Email)
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

func (s *PostgresStore) CreateComment(body, parentId, userId, pageId string) (string, error) {
	id := crypto.Token(21)
	q, args := sb.Insert("comments").
		Columns("id", "body", "parent_id", "user_id", "page_id").
		Values(&id, &body, &parentId, &userId, &pageId).
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

func (s *PostgresStore) GetPage(id string) (*model.Page, error) {
	q, args := sb.Select().From("pages").Where(sb.Eq("id", id)).Query()
	p := &model.Page{}
	err := s.db.QueryRow(q, args...).Scan(&p.Id, &p.Url, &p.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}

	return p, nil
}

func (s *PostgresStore) ListPages(page int, pageSize int) ([]model.Page, error) {
	q, args := sb.Select("id", "url", "title").
		From("pages").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Query()
	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.Page, 0)
	for rows.Next() {
		t := model.Page{}
		if err := rows.Scan(&t.Id, &t.Url, &t.Title); err != nil {
			return nil, err
		}
		out = append(out, t)
	}

	return out, nil
}

func (s *PostgresStore) ListComments(pageId string, pageUrl string, parentId string, page int, pageSize int) ([]model.Comment, error) {
	qb := sb.Select().From("comments")

	if pageId != "" {
		qb = qb.Where(sb.Eq("page_id", pageId))
	} else if pageId == "" && pageUrl != "" {
		qb = qb.Where(sb.Eq("url", pageUrl))
	}
	if parentId != "" && (pageId != "" || pageUrl != "") {
		qb = qb.Where(sb.Eq("parent_id", parentId))
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
		if err := rows.Scan(&o.Id, &o.Body, &o.ParentId, &o.UserId, &o.PageId, &o.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return out, nil
}

func (s *PostgresStore) CountComments(pageId string) (int, error) {
	qb := sb.Select("COUNT(*)").From("comments")
	c := -1
	if pageId != "" {
		q, args := qb.Where(sb.Eq("page_id", pageId)).Query()
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
