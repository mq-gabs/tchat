package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewID() string {
	return strings.Split(uuid.NewString(), "-")[0]
}

func MergeIDs(id1, id2 UserID) (MergedIDs, error) {
	if id1 == "" || id2 == "" {
		return "", errCannotMergeEmptyID
	}
	if id1[0] > id2[0] {
		return MergedIDs(id1 + id2), nil
	}

	return MergedIDs(id2 + id1), nil
}
