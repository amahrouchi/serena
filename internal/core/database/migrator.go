package database

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"gorm.io/gorm"
)

// Migrator is a database migrator (Postgres specific)
type Migrator struct {
	Db *gorm.DB
}

// NewMigrator creates a new Migrator
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		Db: db,
	}
}

// AutoMigrate migrates the schema
func (m *Migrator) AutoMigrate() {
	err := m.Db.AutoMigrate(models.Block{})
	if err != nil {
		panic(err)
	}
}

// ResetDatabase resets the database
func (m *Migrator) ResetDatabase() {
	m.Db.Exec("DROP SCHEMA IF EXISTS public CASCADE")
	m.Db.Exec("CREATE SCHEMA public")

	m.Db.Exec("BEGIN")
	m.Db.Exec("SELECT pg_advisory_xact_lock(12345)")
	m.AutoMigrate()
	m.Db.Exec("COMMIT")
}
