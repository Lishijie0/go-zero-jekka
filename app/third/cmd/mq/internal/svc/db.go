package svc

import (
	"fmt"
	"jekka-api-go/app/third/cmd/mq/internal/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB 创建一个新的数据库连接，并配置连接池（main）
func NewDB(c config.MySQL) (*gorm.DB, error) {
	return newDB(c.DataSource, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxLifetime, logger.Info)
}

// NewDBThird 创建一个新的第三方数据库连接，并配置连接池(third)
func NewDBThird(c config.MySQLThird) (*gorm.DB, error) {
	return newDB(c.DataSource, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxLifetime, logger.Info)
}

// NewDBThirdV2 创建另一个新的第三方数据库连接，并配置连接池(third_v2
func NewDBThirdV2(c config.MySQLThirdV2) (*gorm.DB, error) {
	return newDB(c.DataSource, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxLifetime, logger.Info)
}

// newDB 是一个内部函数，用于创建并配置数据库连接
func newDB(dataSource string, maxIdleConns, maxOpenConns int, connMaxLifetime int, logLevel logger.LogLevel) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层数据库连接失败: %v", err)
	}

	// 验证配置参数
	if err := validateConfig(maxIdleConns, maxOpenConns, connMaxLifetime); err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return db, nil
}

// validateConfig 验证数据库连接池配置参数
func validateConfig(maxIdleConns, maxOpenConns, connMaxLifetime int) error {
	if maxIdleConns < 0 {
		return fmt.Errorf("无效的 MaxIdleConns 值: %d", maxIdleConns)
	}
	if maxOpenConns <= 0 {
		return fmt.Errorf("无效的 MaxOpenConns 值: %d", maxOpenConns)
	}
	if connMaxLifetime <= 0 {
		return fmt.Errorf("无效的 ConnMaxLifetime 值: %d", connMaxLifetime)
	}
	return nil
}
