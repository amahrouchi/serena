package core

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// newDbConnection established the DB connection
func newDbConnection(config *configuration.Config, logger *zerolog.Logger) *gorm.DB {
	logger.Info().Msg("Connecting to the database...")
	db, err := gorm.Open(postgres.Open(config.DbDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Connection to the database established")

	// Migrate the schema
	// TODO: handle migrations properly
	logger.Info().Msg("Automigration of schema: in progress...")
	db.AutoMigrate(models.Block{})
	logger.Info().Msg("Automigration of schema: done!")

	return db
}
