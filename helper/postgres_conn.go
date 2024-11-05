package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func OpenDB(conn string, schm string, ver string) *gorm.DB {
	dsn := conn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   schm + "." + ver + "_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	sqlDB.Close()
	if err != nil {
		panic("Failed to close database")
	}
}
