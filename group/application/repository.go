package application

import (
	d "github.com/bruno5200/CSM/group/domain"
	ds "github.com/bruno5200/CSM/service/domain"
	"github.com/google/uuid"
)

type GroupRepository interface {
	CreateGroup(s *d.Group) error
	ReadGroup(id uuid.UUID) (*d.Group, error)
	ReadGroupsByService(id uuid.UUID) (*[]d.Group, error)
	ReadServiceByKey(key string) (*ds.Service, error)
	UpdateGroup(s *d.Group) error
	DisableGroup(id uuid.UUID) error
}
