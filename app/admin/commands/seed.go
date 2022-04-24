package commands

import (
	"database/sql"
	"fmt"

	"github.com/oussama4/commentify/business/data/store/postgres"
)

type Seed struct {
	db *sql.DB
}

func NewSeed(db *sql.DB) *Seed {
	return &Seed{db: db}
}

func (s *Seed) Synopsis() string {
	return "seed the database with initial data"
}

func (s *Seed) Help() string {
	return "seed the database with initial data"
}

func (s *Seed) Run(args []string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("transaction error %w", err)
	}
	defer tx.Rollback()

	// seed users
	userQuery, userArgs := postgres.UserSeeder()
	userRows, err := tx.Query(userQuery, userArgs...)
	if err != nil {
		return fmt.Errorf("user seeder error %w", err)
	}
	usersIds := []string{}
	for userRows.Next() {
		var userId string
		if err := userRows.Scan(&userId); err != nil {
			return fmt.Errorf("sql scan error %w", err)
		}
		usersIds = append(usersIds, userId)
	}

	// seed pages
	pageQuery, pageArgs := postgres.PageSeeder()
	pageRows, err := tx.Query(pageQuery, pageArgs...)
	if err != nil {
		return fmt.Errorf("page seeder error %w", err)
	}
	pagesIds := []string{}
	for pageRows.Next() {
		var pageId string
		if err := pageRows.Scan(&pageId); err != nil {
			return fmt.Errorf("sql scan error %w", err)
		}
		pagesIds = append(pagesIds, pageId)
	}

	// seed comments
	commentSeeder := postgres.CommentSeeder(usersIds, pagesIds)
	commentQuery, commentArgs := commentSeeder()
	_, err = tx.Exec(commentQuery, commentArgs...)
	if err != nil {
		return fmt.Errorf("comment seeder error %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit error %w", err)
	}

	fmt.Println("databae seeded successfully")
	return nil
}
