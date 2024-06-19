package domain

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	dg "github.com/bruno5200/CSM/group/domain"
	"github.com/bruno5200/CSM/memcache"
	u "github.com/bruno5200/CSM/util"
	"github.com/google/uuid"
)

const (
	API_KEY = "X-API-KEY"
)

var (
	ErrInvalidServiceId   = errors.New("invalid service id")                                                                                                           // invalid service id
	ErrInvalidServiceKey  = errors.New("invalid service key")                                                                                                          // invalid service key
	ErrInvalidServiceName = errors.New("invalid service name, must be at least 3 characters long and contain only lowercase characters and underscores, [a-z_]{3,80}") // invalid service name
	ErrServiceNameEmpty   = errors.New("service name is empty")                                                                                                        // service name is empty
	ErrServiceNotFound    = errors.New("service not found")                                                                                                            // service not found
	ErrNotServices        = errors.New("no services found")                                                                                                            // no services found
	ErrCreatingService    = errors.New("error creating service")                                                                                                       // error creating service
	ErrUpdatingService    = errors.New("error updating service")                                                                                                       // error updating service
	ErrGettingService     = errors.New("error getting service")                                                                                                        // error getting service
	ErrGettingServices    = errors.New("error getting services")                                                                                                       // error getting services
)

func UnmarshalService(data []byte) (*Service, error) {
	var s Service
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s *Service) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

// Item returns an item with the service id as the key
func (s *Service) Item() *memcache.Item {
	data, _ := s.Marshal()
	return &memcache.Item{
		Key:        s.Id.String(),
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()/1000),
	}
}

// ItemName returns an item with the service name as the key
func (s *Service) ItemName() *memcache.Item {
	data, _ := s.Marshal()
	return &memcache.Item{
		Key:        s.Name,
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()/1000),
	}
}

// ItemKey returns an item with the service key as the key
func (s *Service) ItemKey() *memcache.Item {
	data, _ := s.Marshal()
	return &memcache.Item{
		Key:        s.Key,
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()/1000),
	}
}

type Service struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Key         string     `json:"key,omitempty"`
	Description string     `json:"description,omitempty"`
	Groups      []dg.Group `json:"groups,omitempty"`
	Active      bool       `json:"active"`
}

// Update the service with the values from the service request
func (s *Service) Update(sr *ServiceRequest) {
	if sr.Name != "" {
		s.Name = sr.Name
	}

	if sr.Description != "" {
		s.Description = sr.Description
	}
}

// Make new service with the values from the service request
func NewService(sr *ServiceRequest) (*Service, error) {

	if sr.Name == "" {
		return nil, ErrServiceNameEmpty
	}

	if !validName(sr.Name) {
		return nil, ErrInvalidServiceName
	}

	key, err := u.HashPassword(sr.Name)

	if err != nil {
		log.Printf("HASH: %s", err)
		return nil, ErrCreatingService
	}

	u.CheckPasswordHash(sr.Name, key)

	return &Service{
		Id:          uuid.New(),
		Name:        sr.Name,
		Key:         key,
		Description: sr.Description,
		Active:      true,
	}, nil
}

func UnmarshalServiceRequest(data []byte) (*ServiceRequest, error) {
	var sr ServiceRequest
	err := json.Unmarshal(data, &sr)
	return &sr, err
}

type ServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func validName(name string) bool {
	if strings.Contains(name, " ") {
		return false
	}
	if strings.ContainsAny(name, "1234567890") {
		return false
	}
	if strings.ContainsAny(name, "\\|\"≠”¨`´!@$%^*()+=[]{};',.<>?çæ·-:/#") {
		return false
	}
	if strings.ContainsAny(name, "ABCDEFGHIJKLMNOPQRSTUVWXYZÑ") {
		return false
	}
	return (len(name) >= 2 && len(name) <= 80)
}
