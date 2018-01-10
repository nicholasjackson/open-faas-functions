package function

import (
	"net/url"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Handle a serverless request
func Handle(req []byte) string {
	var action string
	query, err := url.ParseQuery(os.Getenv("Http_Query"))
	if err == nil {
		action = query.Get("action")
	}

	if action == "validate" {
		parts := strings.Split(string(req), " ")
		if len(parts) < 2 {
			return "error: decrypt requires two parts separated by a space \"[hash] [password]\""
		}

		if err := bcrypt.CompareHashAndPassword([]byte(parts[0]), []byte(parts[1])); err != nil {
			return "error: not equal"
		}

		return "ok"
	}

	hash, err := bcrypt.GenerateFromPassword(req, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}
