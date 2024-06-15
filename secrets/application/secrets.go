package application

import "github.com/google/uuid"

type Secreter interface {
	GetSecret(secretId uuid.UUID) (string, error)
	CreateSecret(secretId uuid.UUID, secret string) error
}

type Storageer interface {
	ReadSecret(secretId uuid.UUID) (string, error)
	CreateSecret(secretId uuid.UUID, secret string) error
}
