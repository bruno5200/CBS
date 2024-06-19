package application

import (
	"log"

	d "github.com/bruno5200/CSM/block/domain"
	"github.com/bruno5200/CSM/memcache"
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
	Cache     *memcache.Client
}

func (s *blockService) CreateBlock(block *d.Block) error {

	if err := s.Cache.Set(block.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.Cache.Set(block.ItemCheksum()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.BlockRepo.CreateBlock(block); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrCreatingBlock
	}

	return nil
}

func (s *blockService) GetBlock(id uuid.UUID) (*d.Block, error) {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if block, err := d.UnmarshalBlock(item.Value); err == nil {

			if err := s.Cache.Set(block.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			return block, nil

		} else {
			log.Printf("BLOCK: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if block, err := s.BlockRepo.ReadBlock(id); err == nil {

		if err := s.Cache.Set(block.Item()); err != nil {
			log.Printf("MEM: %s", err)
		}

		return block, nil

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingBlock
	}
}

func (s *blockService) GetBlockByCheksum(checksum string) (*d.Block, error) {

	if item, err := s.Cache.Get(checksum); err == nil {

		if block, err := d.UnmarshalBlock(item.Value); err == nil {

			if err := s.Cache.Set(block.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

			if err := s.Cache.Set(block.ItemCheksum()); err != nil {
				log.Printf("MEM: %s", err)
			}

			return block, nil

		} else {
			log.Printf("BLOCK: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if block, err := s.BlockRepo.ReadBlockByCheksum(checksum); err == nil {

		return block, s.Cache.Set(block.Item())

	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingBlock
	}
}

func (s *blockService) GetBlocksByGroup(groupId uuid.UUID) (*[]d.Block, error) {
	if blocks, err := s.BlockRepo.ReadBlocksByGroup(groupId); err == nil {
		return blocks, nil
	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingBlock
	}
}

func (s *blockService) GetBlocksByService(serviceId uuid.UUID) (*[]d.Block, error) {
	if blocks, err := s.BlockRepo.ReadBlocksByService(serviceId); err == nil {
		return blocks, nil
	} else {
		log.Printf("DB: %s", err)
		return nil, d.ErrGettingBlock
	}
}

func (s *blockService) UpdateBlock(block *d.Block) error {

	if err := s.Cache.Set(block.Item()); err != nil {
		log.Printf("MEM: %s", err)
	}

	if err := s.BlockRepo.UpdateBlock(block); err != nil {
		log.Printf("DB: %s", err)
		return d.ErrBlockNotFound
	}

	return nil
}

func (s *blockService) DeleteBlock(id uuid.UUID) error {

	if item, err := s.Cache.Get(id.String()); err == nil {

		if block, err := d.UnmarshalBlock(item.Value); err == nil {

			block.Active = false

			if err := s.Cache.Set(block.Item()); err != nil {
				log.Printf("MEM: %s", err)
			}

		} else {
			log.Printf("BLOCK: %s", err)
		}

	} else {
		log.Printf("MEM: %s", err)
	}

	if err := s.BlockRepo.DisableBlock(id); err != nil {
		log.Printf("DB: %s", err)
		return err
	}

	return nil
}

func NewBlockService(repo BlockRepository, cache *memcache.Client) *blockService {
	return &blockService{BlockRepo: repo, Cache: cache}
}
