package domain

import (
	dg "github.com/bruno5200/CSM/group/domain"
	"github.com/google/uuid"
)

type Service struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Key         string     `json:"kay,omitempty"`
	Description string     `json:"description,omitempty"`
	Groups      []dg.Group `json:"groups,omitempty"`
	Active      bool       `json:"active"`
}
