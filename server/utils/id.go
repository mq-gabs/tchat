package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewID() string {
	return strings.Split(uuid.NewString(), "-")[0]
}

func MergeIDs(id1, id2 string) (Merged2UsersID, error) {
	if id1 == "" || id2 == "" {
		return "", errCannotMergeEmptyID
	}
	if id1[0] > id2[0] {
		return Merged2UsersID(id1 + id2), nil
	}

	return Merged2UsersID(id2 + id2), nil
}
