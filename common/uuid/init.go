package uuid

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct {
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (*UUIDGenerator) GenerateUUID() (string, error) {
	uid, err := uuid.NewRandom()
	return uid.String(), err
}
