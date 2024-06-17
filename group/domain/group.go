package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/bruno5200/CSM/memcache"
	"github.com/google/uuid"
)

var (
	ErrInvalidGroupId        = errors.New("invalid group id")         // invalid group id
	ErrInvalidGroupServiceId = errors.New("invalid group service id") // invalid group service id
	ErrGettingGroup          = errors.New("error getting group")      // error getting group
	ErrGroupNotFound         = errors.New("group not found")          // group not found
	ErrNotGroups             = errors.New("no groups found")          // no groups found
	ErrCreatingGroup         = errors.New("error creating group")     // error creating group
)

func UnmarshalGroup(data []byte) (*Group, error) {
	var g Group
	err := json.Unmarshal(data, &g)
	return &g, err
}

func (g *Group) Marshal() ([]byte, error) {
	return json.Marshal(g)
}

func (g *Group) Item() *memcache.Item {
	data, _ := g.Marshal()
	return &memcache.Item{
		Key:        g.Id.String(),
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()),
	}
}

type Group struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	ServiceId   uuid.UUID `json:"serviceId"`
	ServiceName string    `json:"serviceName,omitempty"`
	Active      bool      `json:"active"`
}

func (g *Group) Update(gr *GroupRequest) {
	if gr.Name != "" {
		g.Name = gr.Name
	}

	if gr.Description != "" {
		g.Description = gr.Description
	}

	serviceId, err := uuid.Parse(gr.ServiceId)

	if err != nil || serviceId == uuid.Nil {
		g.ServiceId = serviceId
	}
}

func NewGroup(gr *GroupRequest) (*Group, error) {

	serviceId, err := uuid.Parse(gr.ServiceId)

	if err != nil || serviceId == uuid.Nil {
		return nil, ErrInvalidGroupServiceId
	}

	return &Group{
		Id:          uuid.New(),
		Name:        gr.Name,
		Description: gr.Description,
		ServiceId:   serviceId,
		Active:      true,
	}, nil
}

func UnmarshalGroupRequest(data []byte) (*GroupRequest, error) {
	var gr GroupRequest
	err := json.Unmarshal(data, &gr)
	return &gr, err
}

type GroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ServiceId   string `json:"serviceId"`
}
