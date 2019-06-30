package utils

import (
	"github.com/google/uuid"
)

// UUID - Generate a new UUID
func UUID() string {
	newUUID := uuid.New()

	return newUUID.String()
}
