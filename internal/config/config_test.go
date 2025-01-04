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

	assert.NotEmpty(t, dbName)
	assert.NotEmpty(t, dbPort)
	assert.NotEmpty(t, dbHost)
	assert.NotEmpty(t, dbPass)
	assert.NotEmpty(t, GetAppEnvLocation())

	assert.Equal(t, "Test_Name", dbName)
	assert.Equal(t, "Test_Pass", dbPass)
	assert.Equal(t, "Test_HOST", dbHost)
	assert.Equal(t, 1234, dbPort)
}
