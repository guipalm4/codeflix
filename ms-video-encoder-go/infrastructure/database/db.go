package database

import (
	"encoder/domain"
	"github.com/jinzhu/gorm"
	"log"
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
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("test db error: %v", err)
	}

	return connection
}

func (db *Database) Connect() (*gorm.DB, error) {

	var err error

	if db.Env != "test" {
		db.Db, err = gorm.Open(db.DbType, db.Dsn)
	} else {
		db.Db, err = gorm.Open(db.DbTypeTest, db.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if db.Debug {
		db.Db.LogMode(true)
	}

	if db.AutoMigrateDb {
		db.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return db.Db, nil
}