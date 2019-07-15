package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorguru/perkbox/types"
)

func TestGetConf(t *testing.T) {
	os.Setenv("DSN_DB", "db_dns")
	os.Setenv("SQL_MIGRATIONS_DIR", "migrations_dir")
	res := getConf()
	expected := types.Environment{
		DSNDB:         "db_dns",
		MigrationsDir: "migrations_dir",
	}
	assert.Equal(t, expected, res)
}
