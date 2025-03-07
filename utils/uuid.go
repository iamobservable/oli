package utils

import (
	"github.com/google/uuid"
)

func UUIDString(data *string) string {
	if data == nil {
		return uuid.New().String()
	} else {
		return uuid.NewSHA1(uuid.NameSpaceURL, []byte(*data)).String()
	}
}
