package store

import (
	"crypto/rand"
	"strings"
)

var alphabet = []byte("ModuleSymbhasOwnPr-0123456789ABCDEFGHNRVfgctiUvz_KqYTJkLxpZXIjQW")

const (
	size = 21
)

func Uid() string {
	bytes := make([]byte, size)
	rand.Read(bytes)

	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteByte(alphabet[bytes[i]&63])
	}

	return b.String()
}
