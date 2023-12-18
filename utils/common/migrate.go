package common

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func Migrateup() {
	m, err := migrate.New(
		"file://db/migration",
		"postgres://postgres:167916@localhost:5432/test_migrate?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	fmt.Println("db migrated successfully")
}
