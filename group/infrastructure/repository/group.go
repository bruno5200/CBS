package repository

import (
	"database/sql"
	"log"

	rb "github.com/bruno5200/CSM/block/infrastructure/repository"
	d "github.com/bruno5200/CSM/group/domain"
	ds "github.com/bruno5200/CSM/service/domain"
	"github.com/bruno5200/CSM/service/infrastructure/repository"
	u "github.com/bruno5200/CSM/util"
	"github.com/google/uuid"
)

const (
	psqlCreateGroup            = `SELECT storage.fn_create_group($1, $2, $3, $4);`
	psqlReadGroup              = `SELECT id, name, service_id, service_name, active FROM storage.fn_read_group($1);`
	psqlReadGroupsByServiceKey = `SELECT id, name, service_id, service_name, active FROM storage.fn_read_groups_by_service_key($1);`
	psqlUpdateGroup            = `SELECT storage.fn_update_group($1, $2, $3, $4);`
	psqlDeleteGroup            = `SELECT storage.fn_delete_group($1);`
)

type groupDB struct {
	Id          uuid.UUID
	Name        sql.NullString
	Description sql.NullString
	ServiceId   uuid.UUID
	ServiceName sql.NullString
	Active      sql.NullBool
}

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *groupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) CreateGroup(g *d.Group) error {

	log.Printf("DB: PSQL, F: storage.fn_create_group('%s', '%s', '%s', '%s'), O:INSERT, T: storage.group", g.Id, g.Name, g.Description, g.ServiceId)
	_, err := r.db.Exec(psqlCreateGroup, g.Id, g.Name, g.Description, g.ServiceId)

	return err
}

func (r *groupRepository) ReadGroup(id uuid.UUID) (*d.Group, error) {

	g := groupDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_group('%s'), O:SELECT, T: storage.group", id)
	err := r.db.QueryRow(psqlReadGroup, id).Scan(&g.Id, &g.Name, &g.ServiceId, &g.ServiceName, &g.Active)

	group := pointerGroup(group(g))

	group.Blocks, _ = rb.NewBlockRepository(r.db).ReadBlocksByGroup(id)

	return group, err
}

func (r *groupRepository) ReadGroupsByService(id uuid.UUID) (*[]d.Group, error) {

	var groups []d.Group

	log.Printf("DB: PSQL, F: storage.fn_read_groups_by_service('%s'), O:SELECT, T: storage.group", id)
	rows, err := r.db.Query(psqlReadGroupsByServiceKey, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		g := groupDB{}
		if err := rows.Scan(&g.Id, &g.Name, &g.ServiceId, &g.ServiceName, &g.Active); err != nil {
			return nil, err
		}
		groups = append(groups, group(g))
	}

	return &groups, nil
}

func (r *groupRepository) ReadServiceByKey(key string) (*ds.Service, error) {
	return repository.NewServiceRepository(r.db).ReadServiceByKey(key)
}

func (r *groupRepository) UpdateGroup(g *d.Group) error {

	log.Printf("DB: PSQL, F: storage.fn_update_group('%s', '%s', '%s', '%s'), O:UPDATE, T: storage.group", g.Id, g.Name, g.Description, g.ServiceId)
	_, err := r.db.Exec(psqlUpdateGroup, g.Id, g.Name, g.Description, g.ServiceId)

	return err
}

func (r *groupRepository) DisableGroup(id uuid.UUID) error {

	log.Printf("DB: PSQL, F: storage.fn_delete_group('%s'), O:UPDATE, T: storage.group", id)
	_, err := r.db.Exec(psqlDeleteGroup, id)

	return err
}

func group(g groupDB) d.Group {
	return d.Group{
		Id:          g.Id,
		Name:        u.NullToString(g.Name),
		Description: u.NullToString(g.Description),
		ServiceId:   g.ServiceId,
		ServiceName: u.NullToString(g.ServiceName),
		Active:      u.NullToBool(g.Active),
	}
}

func pointerGroup(g d.Group) *d.Group {
	return &g
}
