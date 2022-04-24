package commands

import (
	"database/sql"

	"github.com/oussama4/commentify/base/migrate"
)

type Migrate struct {
	db   *sql.DB
	path string
}

func NewMigrateCommand(path string, db *sql.DB) *Migrate {
	return &Migrate{db: db, path: path}
}

func (m *Migrate) Synopsis() string {
	return "migrate the database to the latest version"
}

func (m *Migrate) Help() string {
	return m.Synopsis()
}

func (m *Migrate) Run(args []string) error {
	migrations, err := migrate.Load(m.path)
	if err != nil {
		return err
	}
	migrator := migrate.NewMigrator(migrations, migrate.PostgresDialect{}, m.db)
	if err := migrator.Migrate(); err != nil {
		return err
	}
	return nil
}
