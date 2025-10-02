package db

import (
	"fmt"
	"github/go_auth_api/internal/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBaseService struct {
}

func (db *DataBaseService) Connect(cfg *config.Config) (*gorm.DB, error) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_User,
		cfg.DB_Password,
		cfg.DB_Host,
		cfg.DB_Port,
		cfg.DB_Name,
	)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	gormDB, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})

	fmt.Println("DNS:", dns)

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
