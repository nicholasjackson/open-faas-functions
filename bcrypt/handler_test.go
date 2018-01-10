package function

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"golang.org/x/crypto/bcrypt"
)

func setup(t *testing.T) *is.I {
	return is.New(t)
}

func TestEncrypts(t *testing.T) {
	is := setup(t)
	pwd := []byte("password")

	resp := Handle(pwd)

	err := bcrypt.CompareHashAndPassword([]byte(resp), pwd)
	is.NoErr(err) // Should not have returned an error
}

func TestDecryptsWithSingleItemReturnsError(t *testing.T) {
	is := setup(t)
	defer setEnv("Http_Query", "action=validate")()

	resp := Handle([]byte("abc"))

	is.Equal("error: decrypt requires two parts separated by a space \"[hash] [password]\"", resp) // should have returned an error
}

func TestDecryptsWithInvalidPasswordReturnsError(t *testing.T) {
	is := setup(t)
	defer setEnv("Http_Query", "action=validate")()
	pwd := []byte("password")

	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	resp := Handle([]byte(string(hash) + " tester"))

	is.Equal("error: not equal", resp) // should have returned an error
}

func TestDecrypts(t *testing.T) {
	is := setup(t)
	defer setEnv("Http_Query", "action=validate")()
	pwd := []byte("password")

	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	resp := Handle([]byte(string(hash) + " " + string(pwd)))

	is.Equal("ok", resp) // Should have returned ok
}

func setEnv(key, value string) func() {
	os.Setenv(key, value)

	return func() {
		os.Unsetenv(key)
	}
}
