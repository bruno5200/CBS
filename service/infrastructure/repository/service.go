package repository

import (
	"database/sql"
	"log"

	d "github.com/bruno5200/CSM/service/domain"
	u "github.com/bruno5200/CSM/util"
	"github.com/google/uuid"
)

const (
	psqlCreateService     = `SELECT storage.fn_create_service($1, $2, $3, $4);`
	psqlReadService       = "SELECT id, name, key, description, active FROM storage.fn_read_service($1);"
	psqlReadServiceByName = `SELECT id, name, key, description, active FROM storage.fn_read_service_by_name($1);`
	psqlReadServiceByKey  = `SELECT id, name, key, description, active FROM storage.fn_read_service_by_key($1);`
	psqlReadServices      = `SELECT id, name, key, description, active FROM storage.fn_read_services();`
	psqlUpdateService     = `SELECT storage.fn_update_service($1, $2, $3, $4);`
	psqlDeleteService     = `SELECT storage.fn_disable_service($1);`
)

type serviceDB struct {
	Id          uuid.UUID
	Name        sql.NullString
	Key         sql.NullString
	Description sql.NullString
	Active      sql.NullBool
}

type serviceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *serviceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) CreateService(s *d.Service) error {

	log.Printf("DB: PSQL, F: storage.fn_create_service('%s', '%s', '%s', '%s'), O:INSERT, T: storage.service", s.Id, s.Name, s.Key, s.Description)
	_, err := r.db.Exec(psqlCreateService, s.Id, s.Name, s.Key, s.Description)

	return err
}

func (r *serviceRepository) ReadServices() (*[]d.Service, error) {

	services := []d.Service{}

	log.Printf("DB: PSQL, F: storage.fn_read_services(), O:SELECT, T: storage.service")
	rows, err := r.db.Query(psqlReadServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := serviceDB{}
		if err := rows.Scan(&s.Id, &s.Name, &s.Key, &s.Description, &s.Active); err != nil {
			return nil, err
		}
		services = append(services, service(s))
	}

	return &services, nil
}

func (r *serviceRepository) ReadService(id uuid.UUID) (*d.Service, error) {

	s := serviceDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_service('%s'), O:SELECT, T: storage.service", id)
	err := r.db.QueryRow(psqlReadService, id).Scan(&s.Id, &s.Name, &s.Key, &s.Description, &s.Active)

	return pointerService(service(s)), err
}

func (r *serviceRepository) ReadServiceByName(name string) (*d.Service, error) {

	s := serviceDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_service_by_name('%s'), O:SELECT, T: storage.service", name)
	if err := r.db.QueryRow(psqlReadServiceByName, name).Scan(&s.Id, &s.Name, &s.Description, &s.Active); err != nil {
		return nil, err
	}

	return pointerService(service(s)), nil
}

func (r *serviceRepository) ReadServiceByKey(key string) (*d.Service, error) {

	s := serviceDB{}

	log.Printf("DB: PSQL, F: storage.fn_read_service_by_key('%s'), O:SELECT, T: storage.service", key)
	err := r.db.QueryRow(psqlReadServiceByKey, key).Scan(&s.Id, &s.Name, &s.Description, &s.Active)

	return pointerService(service(s)), err
}

func (r *serviceRepository) UpdateService(s *d.Service) error {

	log.Printf("DB: PSQL, F: storage.fn_update_service('%s', '%s', '%s', '%s'), O:UPDATE, T: storage.service", s.Id, s.Name, s.Key, s.Description)
	_, err := r.db.Exec(psqlUpdateService, s.Id, s.Name, s.Key, s.Description)

	return err
}

func (r *serviceRepository) DisableService(id uuid.UUID) error {

	log.Printf("DB: PSQL, F: storage.fn_disable_service('%s'), O:UPDATE, T: storage.service", id)
	_, err := r.db.Exec(psqlDeleteService, id)

	return err
}

func service(s serviceDB) d.Service {
	return d.Service{
		Id:          s.Id,
		Name:        u.NullToString(s.Name),
		Key:         u.NullToString(s.Key),
		Description: u.NullToString(s.Description),
		Active:      u.NullToBool(s.Active),
	}
}

func pointerService(s d.Service) *d.Service {
	return &s
}
