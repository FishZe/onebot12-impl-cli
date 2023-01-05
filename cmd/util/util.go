package util

import (
	"github.com/gofrs/uuid"
)

func GetUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		return GetUUID()
	}
	return id.String()
}
