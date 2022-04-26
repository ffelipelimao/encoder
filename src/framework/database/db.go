package database

import (
	"log"

	"github.com/ffelipelimao/encoder/src/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "Test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}
	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	if d.Env != "Test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	// TODO: Switch to debug mode
	if d.Debug {
		d.Db.Logger.LogMode(logger.Info)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
