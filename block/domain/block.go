package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/bruno5200/CSM/memcache"
	"github.com/google/uuid"
)

const (
	FilesDir        string = "./files"
	HeaderServiceId string = "X-SVC-ID"
)

var (
	ErrInvalidBlockId        = errors.New("invalid block id")         // invalid block id
	ErrInvalidFileExtension  = errors.New("invalid file extension")   // invalid file extension
	ErrInvalidBlockServiceId = errors.New("invalid block service id") // invalid block service id
	ErrInvalidBlockGroupId   = errors.New("invalid block group id")   // invalid block group id
	ErrMalformedFormKey      = errors.New("form key must be 'file'")  // form key must be 'file'
	ErrBlockNotFound         = errors.New("Block not found")          // Block not found
	ErrNotBlocks             = errors.New("no blocks found")          // no blocks found
	ErrCreatingBlock         = errors.New("error creating block")     // error creating block
	ErrUpdatingBlock         = errors.New("error updating block")     // error updating block
	ErrGettingBlock          = errors.New("error getting block")      // error getting block
)

func UnmarshalBlock(data []byte) (*Block, error) {
	var b Block
	err := json.Unmarshal(data, &b)
	return &b, err
}

func (b *Block) Marshal() ([]byte, error) {
	return json.Marshal(b)
}

// Item returns an item with the block id as the key
func (b *Block) Item() *memcache.Item {
	data, _ := b.Marshal()
	return &memcache.Item{
		Key:        b.Id.String(),
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()/1000),
	}
}

// ItemCheksum returns an item with the checksum as the key
func (b *Block) ItemCheksum() *memcache.Item {
	data, _ := b.Marshal()
	return &memcache.Item{
		Key:        b.Checksum,
		Value:      data,
		Expiration: int32(time.Now().AddDate(0, 1, 0).Unix() - time.Now().Unix()/1000),
	}
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
