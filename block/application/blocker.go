package application

import (
	d "github.com/bruno5200/CSM/block/domain"
	"github.com/google/uuid"
)

type Blocker interface {
	GetBlock(id uuid.UUID) (*d.Block, error)
	GetBlockByCheksum(checksum string) (*d.Block, error)
	GetBlocksByGroup(groupId uuid.UUID) (*[]d.Block, error)
	GetBlocksByService(serviceId uuid.UUID) (*[]d.Block, error)
	CreateBlock(s *d.Block) error
	UpdateBlock(s *d.Block) error
	DeleteBlock(id uuid.UUID) error
}

type blockService struct {
	BlockRepo BlockRepository
}

func (s *blockService) CreateBlock(block *d.Block) error {
	return s.BlockRepo.CreateBlock(block)
}

func (s *blockService) GetBlock(id uuid.UUID) (*d.Block, error) {
	return s.BlockRepo.ReadBlock(id)
}

func (s *blockService) GetBlockByCheksum(checksum string) (*d.Block, error) {
	return s.BlockRepo.ReadBlockByCheksum(checksum)
}

func (s *blockService) GetBlocksByGroup(groupId uuid.UUID) (*[]d.Block, error) {
	return s.BlockRepo.ReadBlocksByGroup(groupId)
}

func (s *blockService) GetBlocksByService(serviceId uuid.UUID) (*[]d.Block, error) {
	return s.BlockRepo.ReadBlocksByService(serviceId)
}

func (s *blockService) UpdateBlock(block *d.Block) error {
	return s.BlockRepo.UpdateBlock(block)
}

func (s *blockService) DeleteBlock(id uuid.UUID) error {
	return s.BlockRepo.DisableBlock(id)
}

func NewBlockService(repo BlockRepository) *blockService {
	return &blockService{BlockRepo: repo}
}
