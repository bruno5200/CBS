package application

import (
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	"github.com/bruno5200/CSM/memcache"
	ds "github.com/bruno5200/CSM/service/domain"
	"github.com/google/uuid"
)

type Grouper interface {
	GetGroup(id uuid.UUID) (*d.Group, error)
	GetGroupsByService(id uuid.UUID) (*[]d.Group, error)
	GetServiceByKey(key string) (*ds.Service, error)
	CreateGroup(s *d.Group) error
	UpdateGroup(s *d.Group) error
	DeleteGroup(id uuid.UUID) error
}

type groupService struct {
	GroupRepo GroupRepository
	Cache     *memcache.Client
}

func (s *groupService) CreateGroup(group *d.Group) error {

	if err := s.Cache.Set(group.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.GroupRepo.CreateGroup(group); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrCreatingGroup
	}

	return nil
}

func (s *groupService) GetGroup(id uuid.UUID) (*d.Group, error) {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if group, err := d.UnmarshalGroup(item.Value); err == nil {

			if err := s.Cache.Set(group.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			return group, nil

		} else {
			log.Printf("GROUP: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if group, err := s.GroupRepo.ReadGroup(id); err == nil {

		if err := s.Cache.Set(group.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return group, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingGroup
	}
}

func (s *groupService) GetGroupsByService(id uuid.UUID) (*[]d.Group, error) {

	if groups, err := s.GroupRepo.ReadGroupsByService(id); err == nil {

		for _, group := range *groups {

			if err := s.Cache.Set(group.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

		}

		return groups, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingGroup
	}
}

func (s *groupService) UpdateGroup(group *d.Group) error {

	if err := s.Cache.Set(group.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.GroupRepo.UpdateGroup(group); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrGroupNotFound
	}

	return nil
}

func (s *groupService) DeleteGroup(id uuid.UUID) error {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if group, err := d.UnmarshalGroup(item.Value); err == nil {

			group.Active = false

			if err := s.Cache.Set(group.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

		} else {
			log.Printf("GROUP: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if err := s.GroupRepo.DisableGroup(id); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrGroupNotFound
	}
	return nil
}

func (s *groupService) GetServiceByKey(key string) (*ds.Service, error) {

	if item, err := s.Cache.Get(key); err == nil {

		if service, err := ds.UnmarshalService(item.Value); err == nil {

			if err := s.Cache.Set(service.Item()); err != nil {
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

	if service, err := s.GroupRepo.ReadServiceByKey(key); err == nil {

		if err := s.Cache.Set(service.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return service, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, ds.ErrGettingService
	}
}

func NewGroupService(repo GroupRepository, cache *memcache.Client) Grouper {
	return &groupService{GroupRepo: repo, Cache: cache}
}
