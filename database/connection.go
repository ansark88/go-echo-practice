package database

import (
	"fmt"
	"go-echo-practice/config"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.DatabaseNew().(*config.DatabaseConfig)
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			databaseConfig.Db.DbHost,
			databaseConfig.Db.DbPort,
			databaseConfig.Db.DbUsername,
			databaseConfig.Db.DbDatabase,
			databaseConfig.Db.DbPassword,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("DB接続エラー :%v", err)
		}
		instance = db
	})
	return instance
}
