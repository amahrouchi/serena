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
	dsn := "host=" + config.App.Db.Host + " user=" + config.App.Db.User + " password=" + config.App.Db.Password + " dbname=" + config.App.Db.DbName + " port=" + config.App.Db.Port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
