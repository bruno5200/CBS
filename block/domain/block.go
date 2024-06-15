package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	FilesDir        = "./files"
	HeaderServiceId = "X-SVC-ID"
)

var (
	ErrInvalidFileExtension = errors.New("invalid file extension")
	ErrInvalidServiceId     = errors.New("invalid service id")
	ErrInvalidGroupId       = errors.New("invalid group id")
	ErrMalformedFormKey     = errors.New("form key must be 'file'")
)

func UnmarshalBlock(data []byte) (*Block, error) {
	var b Block
	err := json.Unmarshal(data, &b)
	return &b, err
}

type Block struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Checksum    string    `json:"checksum,omitempty"`
	Extension   string    `json:"extension"`
	Url         string    `json:"url,omitempty"`
	UploadedAt  time.Time `json:"uploadedAt,omitempty"`
	GroupId     uuid.UUID `json:"groupId"`
	GroupName   string    `json:"groupName,omitempty"`
	ServiceId   uuid.UUID `json:"serviceId"`
	ServiceName string    `json:"serviceName,omitempty"`
	Active      bool      `json:"active"`
}

func (b *Block) Update(block *Block) {
	if block.Name != "" {
		b.Name = block.Name
	}

	if block.GroupId != uuid.Nil {
		b.GroupId = block.GroupId
	}

	if block.ServiceId != uuid.Nil {
		b.ServiceId = block.ServiceId
	}

	if block.Active {
		b.Active = block.Active
	}
}

func NewBlock(name, checksum, url, extension string, id, groupId, serviceId uuid.UUID) *Block {
	return &Block{
		Id:        id,
		Name:      name,
		Checksum:  checksum,
		Url:       url,
		Extension: extension,
		GroupId:   groupId,
		ServiceId: serviceId,
	}
}

func UnmarshalBlockRequest(data []byte) (*BlockRequest, error) {
	var br BlockRequest
	err := json.Unmarshal(data, &br)
	return &br, err
}

type BlockRequest struct {
	Content   string `json:"content"`   // required
	Name      string `json:"name"`      // required
	Extension string `json:"extension"` // required
}
