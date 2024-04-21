package database

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDbConnection established the DB connection
func NewPostgresDbConnection(config *configuration.Config, logger *zerolog.Logger) *gorm.DB {
	logger.Info().Msg("Connecting to the database...")
	dsn := "host=" + config.DbHost + " user=" + config.DbUser + " password=" + config.DbPassword + " dbname=" + config.DbName + " port=" + config.DbPort + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	if config.Env != configuration.EnvTest {
		// TODO: handle migrations properly (test env can keep this behaviour, see RunTestApp func)
		logger.Info().Msg("Auto-migration of the schema...")
		AutoMigrate(db)
	}

	return db
}

// AutoMigrate migrates the schema
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(models.Block{})
	if err != nil {
		panic(err)
	}
}

// ResetDatabase resets the database
func ResetDatabase(db *gorm.DB) {
	db.Exec("DROP SCHEMA IF EXISTS public CASCADE")
	db.Exec("CREATE SCHEMA public")

	// Migrate the schema (Postgres specific)
	db.Exec("BEGIN")
	db.Exec("SELECT pg_advisory_xact_lock(12345)")
	AutoMigrate(db)
	db.Exec("COMMIT")
}
