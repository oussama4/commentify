package migrate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type MigrationFunc func(tx *sql.Tx) error
type Migrations []Migration

func (ms Migrations) Len() int {
	return len(ms)
}

func (ms Migrations) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms Migrations) Less(i, j int) bool {
	timeI := extractMigrationTime(ms[i].Name)
	timeJ := extractMigrationTime(ms[j].Name)
	return timeJ.After(timeI)
}

func extractMigrationTime(migrationName string) time.Time {
	timePart, _, _ := strings.Cut(migrationName, "_")
	t, _ := time.Parse("20060102150405", timePart)
	return t
}

type Migration struct {
	Name string
	Func MigrationFunc
}

func Load(path string) ([]Migration, error) {
	files, err := filepath.Glob(path + "/*.sql")
	if err != nil {
		return nil, fmt.Errorf("error loading migrations %w: ", err)
	}
	var migrations Migrations
	for _, f := range files {
		file, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("File open %w: ", err)
		}
		migration := Migration{
			Name: strings.TrimSuffix(filepath.Base(f), filepath.Ext(f)),
			Func: func(tx *sql.Tx) error {
				_, err := tx.Exec(string(file))
				if err != nil {
					return err
				}
				return nil
			},
		}
		migrations = append(migrations, migration)
	}
	sort.Sort(migrations)
	return migrations, nil
}

func (ms Migrations) Add(migration Migration) {
	ms = append(ms, migration)
}

type Dialect interface {
	CreateTableQuery() string
	GetMigrationQuery() string
	InsertMigrationQuery() string
}

type PostgresDialect struct{}

func (pg PostgresDialect) CreateTableQuery() string {
	return `CREATE TABLE IF NOT EXISTS migrations
		(
		    id		SERIAL PRIMARY KEY,
		    name	CHARACTER VARYING (255) NOT NULL
		)`
}

func (pg PostgresDialect) GetMigrationQuery() string {
	return `SELECT * FROM migrations`
}

func (pg PostgresDialect) InsertMigrationQuery() string {
	return `INSERT INTO migrations (name) VALUES ($1)`
}

type Migrator struct {
	migrations Migrations
	db         *sql.DB
	dialect    Dialect
}

func NewMigrator(migrations Migrations, dialect Dialect, db *sql.DB) *Migrator {
	m := &Migrator{
		migrations: migrations,
		db:         db,
		dialect:    dialect,
	}
	return m
}

func (m *Migrator) appliedMigrations() (Migrations, error) {
	rows, err := m.db.Query(m.dialect.GetMigrationQuery())
	if err != nil {
		return nil, err
	}
	var migrations Migrations
	for rows.Next() {
		id := 0
		migration := Migration{}
		if err := rows.Scan(&id, &migration.Name); err != nil {
			return nil, err
		}
		migrations = append(migrations, migration)
	}
	return migrations, nil
}

func (m *Migrator) pendingMigrations() (Migrations, error) {
	applied, err := m.appliedMigrations()
	if err != nil {
		return nil, err
	}
	if len(applied) == 0 {
		return m.migrations, nil
	}
	sort.Sort(applied)
	pending := Migrations{}
	latestMigrationTime := extractMigrationTime(applied[len(applied)-1].Name)
	for _, migration := range m.migrations {
		migrationTime := extractMigrationTime(migration.Name)
		if migrationTime.After(latestMigrationTime) {
			pending = append(pending, migration)
		}
	}
	return pending, nil
}

func (m *Migrator) insertMigration(tx *sql.Tx, name string) error {
	_, err := tx.Exec(m.dialect.InsertMigrationQuery(), name)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) Migrate() error {
	if err := m.createTable(); err != nil {
		return err
	}

	pendingMigrations, err := m.pendingMigrations()
	if err != nil {
		return err
	}
	for _, migration := range pendingMigrations {
		tx, err := m.db.Begin()
		if err != nil {
			return err
		}
		if err := m.insertMigration(tx, migration.Name); err != nil {
			return err
		}
		if err := migration.Func(tx); err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		log.Println("Migration applied: " + migration.Name)
	}
	return nil
}

func (m *Migrator) createTable() error {
	_, err := m.db.Exec(m.dialect.CreateTableQuery())
	if err != nil {
		return err
	}
	return nil
}

func CreateMigration(path string, name string) error {
	fileName := migrationFileName(name)
	_, err := os.Create(filepath.Join(path, fileName+".sql"))
	if err != nil {
		return err
	}
	return nil
}

func migrationFileName(name string) string {
	return time.Now().Format("20060102150405") + "_" + name
}
