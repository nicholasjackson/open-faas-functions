package function

import (
	"os"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {
	env := ""
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "claim_") {
			env += e + "\n"
		}
	}

	return "Claims:\n" + env
}
