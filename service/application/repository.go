package application

import (
	d "github.com/bruno5200/CSM/service/domain"
	"github.com/google/uuid"
)

type ServiceRepository interface {
	CreateService(s *d.Service) error
	ReadService(id uuid.UUID) (*d.Service, error)
	ReadServiceByName(name string) (*d.Service, error)
	ReadServiceByKey(key string) (*d.Service, error)
	ReadServices() (*[]d.Service, error)
	UpdateService(s *d.Service) error
	DisableService(id uuid.UUID) error
}
