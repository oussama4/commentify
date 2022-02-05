package validate

import (
	"testing"
)

func TestEmailHost(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		// Invalid format.
		{email: "", want: false},
		{email: "email@", want: false},
		{email: "email@x", want: false},
		{email: "email@@example.com", want: false},
		{email: ".email@example.com", want: false},
		{email: "email.@example.com", want: false},
		{email: "email..test@example.com", want: false},
		{email: ".email..test.@example.com", want: false},
		{email: "email@at@example.com", want: false},
		{email: "some whitespace@example.com", want: false},
		{email: "email@whitespace example.com", want: false},
		{email: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", want: false},
		{email: "email@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com", want: false},

		// Unresolvable domain.
		{email: "email+extra@wrong.example.com", want: false},

		// Valid.
		{email: "email@gmail.com", want: true},
		{email: "email.email@gmail.com", want: true},
		{email: "email+extra@example.com", want: true},
		{email: "EMAIL@aol.co.uk", want: true},
		{email: "EMAIL+EXTRA@aol.co.uk", want: true},
	}
	for _, tt := range tests {
		if got := EmailHost(tt.email); got != tt.want {
			t.Errorf("EmailHost() = %v, want %v", got, tt.want)
		}
	}
}
