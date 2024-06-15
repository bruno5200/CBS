package application

import (
	d "github.com/bruno5200/CSM/block/domain"
	"github.com/google/uuid"
)

type BlockRepository interface {
	ReadBlock(id uuid.UUID) (*d.Block, error)
	ReadBlockByCheksum(checksum string) (*d.Block, error)
	ReadBlocksByGroup(groupId uuid.UUID) (*[]d.Block, error)
	ReadBlocksByService(serviceId uuid.UUID) (*[]d.Block, error)
	CreateBlock(b *d.Block) error
	UpdateBlock(b *d.Block) error
	DisableBlock(id uuid.UUID) error
}
