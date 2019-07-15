package main

import (
	"os"

	"github.com/victorguru/perkbox/types"
)

func getConf() types.Environment {
	return types.Environment{
		DSNDB:         os.Getenv("DSN_DB"),
		MigrationsDir: os.Getenv("SQL_MIGRATIONS_DIR"),
	}
}
