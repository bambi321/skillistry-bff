package helpers

import "github.com/google/uuid"

func GenerateUUID(s string) uuid.UUID {
	namespace := uuid.NameSpaceURL
	return uuid.NewSHA1(namespace, []byte(s))
}
