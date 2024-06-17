package application

import (
	"log"

	d "github.com/bruno5200/CSM/group/domain"
	"github.com/bruno5200/CSM/memcache"
	"github.com/google/uuid"
)

type Grouper interface {
	GetGroup(id uuid.UUID) (*d.Group, error)
	GetGroupsByService() (*[]d.Group, error)
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

func (s *groupService) GetGroupsByService() (*[]d.Group, error) {
	return s.GroupRepo.ReadGroupsByService()
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

func NewGroupService(repo GroupRepository, cache *memcache.Client) Grouper {
	return &groupService{GroupRepo: repo, Cache: cache}
}
