package application

import d "github.com/bruno5200/CSM/service/domain"

type ServiceRepository interface {
	ReadService(id string) (d.Service, error)
	ReadServices() ([]d.Service, error)
	CreateService(s d.Service) error
	UpdateService(s d.Service) error
	DeleteService(id string) error
}
