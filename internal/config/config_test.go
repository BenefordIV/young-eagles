package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
	LoadConfig("../..")
}

func TestGetConfig(t *testing.T) {
	setup()
	dbName := GetDbName()
	dbPass := GetDbPass()
	dbHost := GetDbHost()
	dbPort := GetDbPort()
	dbUser := GetDbUsername()

	assert.NotEmpty(t, dbName)
	assert.NotEmpty(t, dbPort)
	assert.NotEmpty(t, dbHost)
	assert.NotEmpty(t, dbPass)
	assert.NotEmpty(t, GetAppEnvLocation())

	assert.Equal(t, "young_eagles_developer", dbName)
	assert.Equal(t, "YoungEagles!", dbPass)
	assert.Equal(t, "young-eagle-developer.cxwcc8oqmytb.us-east-1.rds.amazonaws.com", dbHost)
	assert.Equal(t, 5432, dbPort)
	assert.Equal(t, "ye_admin", dbUser)
}
