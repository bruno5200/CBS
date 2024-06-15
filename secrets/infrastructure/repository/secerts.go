package repository

import (
	"database/sql"

	"github.com/bruno5200/CSM/memcache"
	"github.com/bruno5200/CSM/secrets/application"
	"github.com/google/uuid"
)

type secretsDB struct {
	db    *sql.DB
	cache *memcache.Client
}

func NewSecretsDB(db *sql.DB, cache *memcache.Client) application.Storageer {
	return &secretsDB{
		db:    db,
		cache: cache,
	}
}

func (s *secretsDB) ReadSecret(secretId uuid.UUID) (string, error) {
	return "", nil
}

func (s *secretsDB) CreateSecret(secretId uuid.UUID, secret string) error {
	return nil
}
