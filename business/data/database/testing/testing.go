package testing

import (
	"database/sql"
	"log"
	"os"
	"os/exec"

	"github.com/oussama4/commentify/business/data/database"
)

func migrateDatabase() error {
	c := exec.Command("goose",
		"-dir",
		"../../schema",
		"postgres",
		"user=commentify dbname=commentify_testing password=secret sslmode=disable",
		"up",
	)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}
	log.Println("test database migrated")
	return nil
}

func seedDatabase() error {
	c := exec.Command("go",
		"run",
		"../../../../app/admin/main.go",
		"--db-name commentify_testing",
		"seed",
	)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}
	log.Println("test database seeded")
	return nil
}

func SetupDB() *sql.DB {
	if err := migrateDatabase(); err != nil {
		log.Fatal("failed to migrate test database ", err)
	}
	if err := seedDatabase(); err != nil {
		log.Fatal("failed to seed test database ", err)
	}
	dbCfg := database.Config{
		User:       "commentify",
		Password:   "secret",
		Host:       "localhost",
		Name:       "commentify_testing",
		DisableTLS: true,
	}
	db, err := database.Open(dbCfg)
	if err != nil {
		log.Fatal("failed to open a connection to test database", err)
	}

	return db
}

func CleanDB(db *sql.DB) {
	log.Println("cleaning testing database ...")
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		log.Fatalln("error quering database tables ", err)
	}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalln("error scanning rows ", err)
		}
		_, err := db.Exec("DROP TABLE IF EXISTS " + tableName + " CASCADE")
		if err != nil {
			log.Fatalln("error dropping table ", err)
		}
	}
	log.Println("testing database cleaned")
}
