package common

import (
	"roomate/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

type DbMigration interface {
	RunDBMigration()
}

type dbMigration struct {
	cfg config.MigrateConfig
}

func (d *dbMigration) RunDBMigration() {
	migration, err := migrate.New(d.cfg.MigrationUrl, d.cfg.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func NewDbMigration(cfg config.MigrateConfig) DbMigration {
	return &dbMigration{
		cfg: cfg,
	}
}
