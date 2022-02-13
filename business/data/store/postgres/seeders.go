package postgres

import (
	"math/rand"
	"time"

	"github.com/oussama4/commentify/base/faker"
	sb "github.com/oussama4/sqlbuilder"
)

type SeederFunc func() (string, []interface{})

func UserSeeder() (string, []interface{}) {
	b := sb.Insert("users").Columns("id", "name", "email")
	for i := 0; i < 10; i++ {
		b = b.Values(faker.UniqueString(12), faker.UserName(), faker.Email())
	}
	b = b.Returning("id")
	return b.Query()
}

func PageSeeder() (string, []interface{}) {
	b := sb.Insert("pages").Columns("id", "url", "title")
	for i := 0; i < 10; i++ {
		b = b.Values(faker.UniqueString(12), faker.URL(), faker.UserName())
	}
	b = b.Returning("id")
	return b.Query()
}

func CommentSeeder(userIds []string, pageIds []string) SeederFunc {
	sf := func() (string, []interface{}) {
		rand.Seed(time.Now().UnixNano())
		b := sb.Insert("comments").Columns("id", "body", "user_id", "page_id")
		for i := 0; i < 10; i++ {
			userId := userIds[rand.Int()%len(userIds)]
			pageId := pageIds[rand.Int()%len(pageIds)]
			b = b.Values(faker.UniqueString(12), faker.Paragraph(), userId, pageId)
		}
		return b.Query()
	}
	return sf
}
