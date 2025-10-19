package utils

import "testing"

func Test_ID(t *testing.T) {
	id := NewID()

	t.Log(id)

	if id == "" {
		t.Fatal("id is empty")
	}
}
