package postgres

import (
	"database/sql"
	"testing"

	"github.com/oussama4/commentify/base/faker"
	dbtest "github.com/oussama4/commentify/business/data/database/testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	db = dbtest.SetupDB()
	defer db.Close()
	defer dbtest.CleanDB(db)
	m.Run()
}

func TestPostgresStore_CreatePage(t *testing.T) {
	store, err := Create(db)
	if err != nil {
		t.Log(err)
	}

	pageId, err := store.CreatePage(faker.URL(), faker.Word())
	if err != nil {
		t.Log(err)
	}
	page, err := store.GetPage(pageId)
	if err != nil {
		t.Log(err)
	}

	if page.Id != pageId {
		t.Errorf("expected created page id to be %s but got %s", pageId, page.Id)
	}
}
