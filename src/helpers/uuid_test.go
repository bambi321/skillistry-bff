package helpers

import (
	"testing"
)

func TestGenerateUUIDHelloWorld(t *testing.T) {
	test := "hello-world"
	expected := "62e71b4a-dc1e-5b38-96f7-7d398968d40b"
	t.Run(test, func(t *testing.T) {
		result := GenerateUUID(test).String()
		if result != expected {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})
}

func TestGenerateUUIDIsConsistent(t *testing.T) {
	input := "consistency-check"
	u1 := GenerateUUID(input)
	u2 := GenerateUUID(input)

	if u1 != u2 {
		t.Errorf("UUIDs should match for same input. Got %s and %s", u1, u2)
	}
}
