package database

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDbConnection established the DB connection
func NewPostgresDbConnection(
	config *configuration.Config,
	logger *zerolog.Logger,
) *gorm.DB {
	logger.Info().Msg("Connecting to the database...")
	dsn := "host=" + config.DbHost + " user=" + config.DbUser + " password=" + config.DbPassword + " dbname=" + config.DbName + " port=" + config.DbPort + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
