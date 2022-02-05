// faker is a simple package for generating fake data
package faker

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var tlds = []string{"com", "net", "org", "io", "info"}

var wordList = []string{
	"alias", "consequatur", "aut", "perferendis", "sit", "voluptatem",
	"accusantium", "doloremque", "aperiam", "eaque", "ipsa", "quae", "ab",
	"illo", "inventore", "veritatis", "et", "quasi", "architecto",
	"beatae", "vitae", "dicta", "sunt", "explicabo", "aspernatur", "aut",
	"odit", "aut", "fugit", "sed", "quia", "consequuntur", "magni",
	"dolores", "eos", "qui", "ratione", "voluptatem", "sequi", "nesciunt",
	"neque", "dolorem", "ipsum", "quia", "dolor", "sit", "amet",
	"consectetur", "adipisci", "velit", "sed", "quia", "non", "numquam",
	"eius", "modi", "tempora", "incidunt", "ut", "labore", "et", "dolore",
	"magnam", "aliquam", "quaerat", "voluptatem", "ut", "enim", "ad",
	"minima", "veniam", "quis", "nostrum", "exercitationem", "ullam",
	"corporis", "nemo", "enim", "ipsam", "voluptatem", "quia", "voluptas",
	"sit", "suscipit", "laboriosam", "nisi", "ut", "aliquid", "ex", "ea",
	"commodi", "consequatur", "quis", "autem", "vel", "eum", "iure",
	"reprehenderit", "qui", "in", "ea", "voluptate", "velit", "esse",
	"quam", "nihil", "molestiae", "et", "iusto", "odio", "dignissimos",
	"ducimus", "qui", "blanditiis", "praesentium", "laudantium", "totam",
	"rem", "voluptatum", "deleniti", "atque", "corrupti", "quos",
	"dolores", "et", "quas", "molestias", "excepturi", "sint",
	"occaecati", "cupiditate", "non", "provident", "sed", "ut",
	"perspiciatis", "unde", "omnis", "iste", "natus", "error",
	"similique", "sunt", "in", "culpa", "qui", "officia", "deserunt",
	"mollitia", "animi", "id", "est", "laborum", "et", "dolorum", "fuga",
	"et", "harum", "quidem", "rerum", "facilis", "est", "et", "expedita",
	"distinctio", "nam", "libero", "tempore", "cum", "soluta", "nobis",
	"est", "eligendi", "optio", "cumque", "nihil", "impedit", "quo",
	"porro", "quisquam", "est", "qui", "minus", "id", "quod", "maxime",
	"placeat", "facere", "possimus", "omnis", "voluptas", "assumenda",
	"est", "omnis", "dolor", "repellendus", "temporibus", "autem",
	"quibusdam", "et", "aut", "consequatur", "vel", "illum", "qui",
	"dolorem", "eum", "fugiat", "quo", "voluptas", "nulla", "pariatur",
	"at", "vero", "eos", "et", "accusamus", "officiis", "debitis", "aut",
	"rerum", "necessitatibus", "saepe", "eveniet", "ut", "et",
	"voluptates", "repudiandae", "sint", "et", "molestiae", "non",
	"recusandae", "itaque", "earum", "rerum", "hic", "tenetur", "a",
	"sapiente", "delectus", "ut", "aut", "reiciendis", "voluptatibus",
	"maiores", "doloribus", "asperiores", "repellat",
}

func Word() string {
	return randomStringSliceElement(wordList)
}

func Sentence() string {
	s := ""
	count := rand.Intn(10) + 1
	for i := 0; i < count; i++ {
		s += randomStringSliceElement(wordList)
	}
	return fmt.Sprintf("%s.", s)
}

func Paragraph() string {
	p := ""
	count := rand.Intn(10) + 1
	for i := 0; i < count; i++ {
		p += Sentence()
	}
	return p
}

func randomStringSliceElement(s []string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Int()%len(s)]
}

func randomString(size int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, size)
	for k := range b {
		b[k] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func Email() string {
	return randomString(7) + "@" + randomString(7) + "." + randomStringSliceElement(tlds)
}

func UniqueString(size int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, size)
	crand.Read(bytes)

	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteByte(alphabet[bytes[i]%61])
	}

	return b.String()
}

func UserName() string {
	return randomString(7)
}

func DomainName() string {
	return fmt.Sprintf("%s.%s", randomString(7), randomStringSliceElement(tlds))
}

func URL() string {
	return fmt.Sprintf("https://%s/%s", DomainName(), randomString(12))
}
