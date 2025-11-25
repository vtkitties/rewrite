package handlers

import (
	"fmt"
)

func ErrorResponse(message string) string {
	return fmt.Sprint(map[string]any{"error": message})
}
