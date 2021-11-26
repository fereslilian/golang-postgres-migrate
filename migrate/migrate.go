package migrate

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

const defaultMigrationsPath = "file://db/migrations"

func ExecuteMigration(config DBConfig, useSSL bool, migrationsPath string) error {

	validationError := validateDBConfig(config)
	if validationError != nil {
		return validationError
	}

	driver, databaseConnectionError := connectToDatabase(config, useSSL)
	if databaseConnectionError != nil {
		return databaseConnectionError
	}

	migrationError := migrateDatabase(driver, migrationsPath)
	if migrationError != nil {
		return migrationError
	}
	return nil
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func getSSLMode(useSSL bool) string {
	if useSSL {
		return "require"
	} else {
		return "disable"
	}
}

func validateDBConfig(config DBConfig) error {
	if strings.TrimSpace(config.Host) == "" || strings.TrimSpace(config.Port) == "" || strings.TrimSpace(config.Name) == "" || strings.TrimSpace(config.User) == "" || strings.TrimSpace(config.Name) == "" {
		return errors.New("Migrate: DBConfig inválida.")
	}
	return nil
}

func migrateDatabase(driver database.Driver, migrationsPath string) error {

	if len(migrationsPath) == 0 {
		migrationsPath = defaultMigrationsPath
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return errors.New(fmt.Sprintf("Migrate: erro ao ler migrations [%s].", err.Error()))
	}

	if err := m.Up(); err != nil {
		if errors.Cause(err) == migrate.ErrNoChange {
			logger.Info("Migrate: não há novas migrations a serem realizadas")
			return nil
		}
		return errors.New(fmt.Sprintf("Migrate: erro ao realizar migrations [%s].", err.Error()))
	}
	logger.Info("Migrate: sucesso ao realizar migrations")
	return nil
}

func connectToDatabase(config DBConfig, useSSL bool) (database.Driver, error) {
	SSLMode := getSSLMode(useSSL)

	db, err := sql.Open("postgres", "host=" + config.Host + " user=" + config.User + " dbname=" + config.Name + " password=" + config.Password + " sslmode=" + SSLMode)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Migrate: erro ao abrir conexão com database [%s].", err.Error()))
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Migrate: erro ao conectar com database [%s].", err.Error()))
	}
	return driver, nil
}
