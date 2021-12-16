package server

import (
	"strings"
)

func extractBearer(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) != 2 {
		return ""
	}
	return strings.TrimSpace(splitToken[1])
}
