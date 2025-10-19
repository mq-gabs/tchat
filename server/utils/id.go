package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewID() string {
	return strings.Split(uuid.NewString(), "-")[0]
}
