package application

import (
	d "github.com/bruno5200/CSM/group/domain"
	"github.com/google/uuid"
)

type Grouper interface {
	GetGroup(id uuid.UUID) (d.Group, error)
	GetGroupsByService() ([]d.Group, error)
	CreateGroup(s d.Group) error
	UpdateGroup(s d.Group) error
	DeleteGroup(id uuid.UUID) error
}
