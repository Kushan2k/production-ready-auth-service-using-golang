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
	cfg *config.Config
}

func NewDatabaseService(cfg *config.Config) *DataBaseService {
	return &DataBaseService{cfg: cfg}
}

func (db *DataBaseService) Connect() (*gorm.DB, error) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.cfg.DB_User,
		db.cfg.DB_Password,
		db.cfg.DB_Host,
		db.cfg.DB_Port,
		db.cfg.DB_Name,
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
