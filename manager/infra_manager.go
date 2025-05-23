package manager

import (
	"database/sql"
	"fmt"
	"roomate/config"
)

type InfraManager interface {
	openConn() error
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) openConn() error {
	// dsn = data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.Username, i.cfg.Password, i.cfg.DbName)

	db, err := sql.Open(i.cfg.Driver, dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}
