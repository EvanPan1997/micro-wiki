package utils

import uuid "github.com/satori/go.uuid"

func GenUuidV4() string {
	return uuid.NewV4().String()
}
