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
		"postgres://root:167916@postgres:5432/roomate?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	fmt.Println("db migrated successfully")
}
