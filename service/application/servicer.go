package application

import d "github.com/bruno5200/CSM/service/domain"

type Serviceer interface {
	GetService(id string) (d.Service, error)
	GetServices() ([]d.Service, error)
	CreateService(s d.Service) error
	UpdateService(s d.Service) error
	DeleteService(id string) error
}
