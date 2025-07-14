package dependencies

import (
	"fmt"
	"github.com/ghn-rs/cloud-strife-user/internal/dependencies/config"
	"github.com/ghn-rs/cloud-strife-user/internal/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

type DatabaseClient struct {
	Db *gorm.DB
}

func NewDatabaseClient() (*DatabaseClient, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DatabaseClient{Db: db}, nil
}

func (client *DatabaseClient) Close() error {
	sqlDB, err := client.Db.DB()
	if err != nil {
		return fmt.Errorf("failed to get generic database object: %w", err)
	}
	return sqlDB.Close()
}

func (client *DatabaseClient) RunMigrations() error {
	log.Println("Running database migrations")
	err := client.Db.AutoMigrate(&entity.User{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
