package core

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbConnection established the DB connection
func dbConnection(config *Config, logger *zerolog.Logger) {
	db, err := gorm.Open(postgres.Open(config.DbDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Connection to the database established")

	// Migrate the schema
	db.AutoMigrate(models.Block{})
}
