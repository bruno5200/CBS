package application

import (
	"encoding/base64"
	"log"

	"github.com/bruno5200/CSM/memcache"
	d "github.com/bruno5200/CSM/service/domain"
	"github.com/google/uuid"
)

type Serviceer interface {
	GetService(id uuid.UUID) (*d.Service, error)
	GetServiceByName(name string) (*d.Service, error)
	GetServiceByKey(key string) (*d.Service, error)
	GetServices() (*[]d.Service, error)
	CreateService(s *d.Service) error
	UpdateService(s *d.Service) error
	DeleteService(id uuid.UUID) error
}

type serviceService struct {
	ServiceRepo ServiceRepository
	Cache       *memcache.Client
}

func (s *serviceService) CreateService(service *d.Service) error {

	if err := s.Cache.Set(service.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.Cache.Set(service.ItemName()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.Cache.Set(service.ItemKey()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.ServiceRepo.CreateService(service); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrCreatingService
	}

	return nil
}

func (s *serviceService) GetService(id uuid.UUID) (*d.Service, error) {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if service, err := d.UnmarshalService(item.Value); err == nil {

			if err := s.Cache.Set(service.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemName()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemKey()); err != nil {
				log.Printf("MEM: %s", err)
			}

			return service, nil

		} else {
			log.Printf("SERVICE: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if service, err := s.ServiceRepo.ReadService(id); err == nil {

		if err := s.Cache.Set(service.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return service, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingService
	}
}

func (s *serviceService) GetServiceByName(name string) (*d.Service, error) {

	if service, err := s.ServiceRepo.ReadServiceByName(name); err == nil {

		if err := s.Cache.Set(service.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		if err := s.Cache.Set(service.ItemName()); err != nil {
			log.Printf("MEM: %s", err)
		}

		if err := s.Cache.Set(service.ItemKey()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return service, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingService
	}
}

func (s *serviceService) GetServiceByKey(key string) (*d.Service, error) {

	if item, err := s.Cache.Get(key); err == nil {

		if service, err := d.UnmarshalService(item.Value); err == nil {

			if err := s.Cache.Set(service.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemName()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemKey()); err != nil {
				log.Printf("MEM: %s", err)
			}

			return service, nil

		} else {
			log.Printf("SERVICE: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if service, err := s.ServiceRepo.ReadServiceByKey(key); err == nil {

		if err := s.Cache.Set(service.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return service, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingService
	}
}

func (s *serviceService) GetServices() (*[]d.Service, error) {

	if services, err := s.ServiceRepo.ReadServices(); err == nil {

		for _, service := range *services {

			if err := s.Cache.Set(service.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemName()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemKey()); err != nil {
				log.Printf("MEM: %s", err)
			}

			service.Key = base64.StdEncoding.EncodeToString([]byte(service.Key))

		}
		return services, nil
		
	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingServices
	}
}

func (s *serviceService) UpdateService(service *d.Service) error {

	if err := s.Cache.Set(service.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.Cache.Set(service.ItemName()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.Cache.Set(service.ItemKey()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.ServiceRepo.UpdateService(service); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrUpdatingService
	}

	return nil
}

func (s *serviceService) DeleteService(id uuid.UUID) error {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if service, err := d.UnmarshalService(item.Value); err == nil {

			service.Active = false

			if err := s.Cache.Set(service.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemName()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(service.ItemKey()); err != nil {
				log.Printf("MEM: %s", err)
			}

		} else {
			log.Printf("SERVICE: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if err := s.ServiceRepo.DisableService(id); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrServiceNotFound
	}

	return nil
}

func NewServiceService(sr ServiceRepository, cache *memcache.Client) Serviceer {
	return &serviceService{ServiceRepo: sr, Cache: cache}
}
