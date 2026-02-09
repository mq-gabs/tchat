package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewID() string {
	return strings.Split(uuid.NewString(), "-")[0]
}

func MakeChatID(id1, id2 UserID) (ChatID, error) {
	if id1 == "" || id2 == "" {
		return "", errCannotMakeChatID
	}
	if id1[0] > id2[0] {
		return ChatID(id1 + id2), nil
	}

	return ChatID(id2 + id1), nil
}
