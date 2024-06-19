package repository

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/bruno5200/CSM/block/domain"
	"github.com/bruno5200/CSM/env"
	u "github.com/bruno5200/CSM/util"
	"github.com/google/uuid"
)

var e = env.Env()

const (
	psqlCreateBlock         = `SELECT storage.fn_create_block($1, $2, $3, $4, $5, $6);`
	psqlReadBlock           = `SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_block($1);`
	psqlReadBlockByCheksum  = `SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_block_by_checksum($1);`
	psqlReadBlocksByGroup   = `SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_blocks_by_group($1);`
	psqlReadBlocksByService = `SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_blocks_by_service($1);`
	psqlUpdateBlock         = `SELECT storage.fn_update_block($1, $2, $3, $4, $5, $6, $7);`
	psqlDeleteBlock         = `SELECT storage.fn_disable_block($1);`
)

type blockDB struct {
	Id          uuid.UUID
	Name        sql.NullString
	Checksum    sql.NullString
	Extension   sql.NullString
	Url         sql.NullString
	UploadedAt  sql.NullTime
	GroupId     uuid.UUID
	GroupName   sql.NullString
	ServiceId   uuid.UUID
	ServiceName sql.NullString
	Active      sql.NullBool
}

type blockRepository struct {
	db *sql.DB
}

func NewBlockRepository(db *sql.DB) *blockRepository {
	return &blockRepository{db: db}
}

func (r *blockRepository) CreateBlock(b *d.Block) error {

	log.Printf("DB: PSQL, F: storage.fn_create_block('%s', '%s', '%s', '%s', '%s', '%s'), O:INSERT, T: storage.block", b.Id, b.Name, b.Checksum, b.Extension, b.Url, b.GroupId)
	_, err := r.db.Exec(psqlCreateBlock, b.Id, b.Name, b.Checksum, b.Extension, b.Url, b.GroupId)

	return err

}

func (r *blockRepository) ReadBlock(id uuid.UUID) (*d.Block, error) {

	b := blockDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_block('%s'), O:SELECT, T: storage.block", id)
	err := r.db.QueryRow(psqlReadBlock, id).Scan(&b.Id, &b.Name, &b.Checksum, &b.Extension, &b.Url, &b.UploadedAt, &b.GroupId, &b.GroupName, &b.ServiceId, &b.ServiceName, &b.Active)

	return pointerBlock(block(b)), err
}

func (r *blockRepository) ReadBlockByCheksum(checksum string) (*d.Block, error) {

	b := blockDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_block_by_checksum('%s'), O:SELECT, T: storage.block", checksum)
	err := r.db.QueryRow(psqlReadBlockByCheksum, checksum).Scan(&b.Id, &b.Name, &b.Checksum, &b.Extension, &b.Url, &b.UploadedAt, &b.GroupId, &b.GroupName, &b.ServiceId, &b.ServiceName, &b.Active)

	return pointerBlock(block(b)), err
}

func (r *blockRepository) ReadBlocksByGroup(groupId uuid.UUID) (*[]d.Block, error) {
	var blocks []d.Block

	log.Printf("DB: PSQL, F: storage.fn_read_blocks_by_group('%s'), O:SELECT, T: storage.block", groupId)
	rows, err := r.db.Query(psqlReadBlocksByGroup, groupId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b := blockDB{}
		err = rows.Scan(&b.Id, &b.Name, &b.Checksum, &b.Extension, &b.Url, &b.UploadedAt, &b.GroupId, &b.GroupName, &b.ServiceId, &b.ServiceName, &b.Active)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block(b))
	}

	return &blocks, nil
}

func (r *blockRepository) ReadBlocksByService(serviceId uuid.UUID) (*[]d.Block, error) {
	var blocks []d.Block

	log.Printf("DB: PSQL, F: storage.fn_read_blocks_by_service('%s'), O:SELECT, T: storage.block", serviceId)
	rows, err := r.db.Query(psqlReadBlocksByService, serviceId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b := blockDB{}
		err = rows.Scan(&b.Id, &b.Name, &b.Checksum, &b.Extension, &b.Url, &b.UploadedAt, &b.GroupId, &b.GroupName, &b.ServiceId, &b.ServiceName, &b.Active)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block(b))
	}

	return &blocks, nil
}

func (r *blockRepository) UpdateBlock(b *d.Block) error {

	log.Printf("DB: PSQL, F: storage.fn_update_block('%s', '%s', '%s', '%s', '%s', '%s', %t), O:UPDATE, T: storage.block", b.Id, b.Name, b.Checksum, b.Extension, b.Url, b.GroupId, b.Active)
	_, err := r.db.Exec(psqlUpdateBlock, b.Id, b.Name, b.Checksum, b.Extension, b.Url, b.GroupId, b.Active)

	return err
}

func (r *blockRepository) DisableBlock(id uuid.UUID) error {

	log.Printf("DB: PSQL, F: storage.fn_delete_block('%s'), O:DELETE, T: storage.block", id)
	_, err := r.db.Exec(psqlDeleteBlock, id)

	return err
}

func block(b blockDB) d.Block {
	return d.Block{
		Id:          b.Id,
		Name:        u.NullToString(b.Name),
		Checksum:    u.NullToString(b.Checksum),
		Extension:   u.NullToString(b.Extension),
		Url:         fmt.Sprintf("%s/api/v1/block/%s", e.GetUrl(), b.Id),
		UploadedAt:  u.NullToTime(b.UploadedAt),
		GroupId:     b.GroupId,
		GroupName:   u.NullToString(b.GroupName),
		ServiceId:   b.ServiceId,
		ServiceName: u.NullToString(b.ServiceName),
		Active:      u.NullToBool(b.Active),
	}
}

func pointerBlock(b d.Block) *d.Block {
	return &b
}
