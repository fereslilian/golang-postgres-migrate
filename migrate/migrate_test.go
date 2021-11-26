package migrate

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConfigWithInvalidHost(t *testing.T) {

	config := DBConfig{
		Port:     "5432",
		User:     "dbuser",
		Name:     "dbname",
		Password: "23$#%8",
	}

	error := ExecuteMigration(config, false, defaultMigrationsPath)

	assert.Equal(t, errors.New("Migrate: DBConfig inválida."), error)

}

func TestDBConfigWithInvalidPort(t *testing.T) {

	config := DBConfig{
		Host:     "localhost",
		User:     "dbuser",
		Name:     "dbname",
		Password: "23$#%8",
	}

	error := ExecuteMigration(config, false, defaultMigrationsPath)

	assert.Equal(t, errors.New("Migrate: DBConfig inválida."), error)

}

func TestDBConfigWithInvalidUser(t *testing.T) {

	config := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Name:     "dbname",
		Password: "23$#%8",
	}

	error := ExecuteMigration(config, false, defaultMigrationsPath)

	assert.Equal(t, errors.New("Migrate: DBConfig inválida."), error)
}

func TestDBConfigWithInvalidPassword(t *testing.T) {

	config := DBConfig{
		Host: "localhost",
		Port: "5432",
		User: "dbuser",
		Name: "dbname",
	}

	error := ExecuteMigration(config, false, defaultMigrationsPath)

	assert.Equal(t, errors.New("Migrate: DBConfig inválida."), error)
}

func TestDBConfigWithInvalidDatabaseName(t *testing.T) {

	config := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "dbuser",
		Password: "23$#%8",
	}

	error := ExecuteMigration(config, false, defaultMigrationsPath)

	assert.Equal(t, errors.New("Migrate: DBConfig inválida."), error)
}


