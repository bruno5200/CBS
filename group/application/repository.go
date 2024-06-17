package application

import (
	d "github.com/bruno5200/CSM/group/domain"
	"github.com/google/uuid"
)

type GroupRepository interface {
	ReadGroup(id uuid.UUID) (*d.Group, error)
	ReadGroupsByService() (*[]d.Group, error)
	CreateGroup(s *d.Group) error
	UpdateGroup(s *d.Group) error
	DisableGroup(id uuid.UUID) error
}
