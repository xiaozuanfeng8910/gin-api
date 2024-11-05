package db

import (
	"gin-api/internal/config"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

func getGormLogWriter(cfg *config.Config) logger.Writer {
	dbConfig := cfg.Database["ci_db"]
	logConfig := &cfg.Log
	var writer io.Writer
	//是否启用日志文件
	if dbConfig.EnableFileLogWriter {
		// 自定义writer
		writer = &lumberjack.Logger{
			Filename:   logConfig.RootDir + "/" + dbConfig.LogFilename,
			MaxSize:    logConfig.MaxSize,
			MaxBackups: logConfig.MaxBackups,
			MaxAge:     logConfig.MaxAge,
			Compress:   logConfig.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger(cfg *config.Config) logger.Interface {
	dbConfig := cfg.Database["ci_db"]
	var logMode logger.LogLevel
	switch dbConfig.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(cfg), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢SQL阈值
		LogLevel:                  logMode,                // 日志级别
		IgnoreRecordNotFoundError: false,
		Colorful:                  !dbConfig.EnableFileLogWriter, // 禁用彩色打印
	})
}

// InitMySQLGorm 初始化MySQL gorm.DB
func InitMySQLGorm(cfg *config.Config, zapLogger *zap.Logger) (*gorm.DB, error) {
	dbConfig := cfg.Database["ci_db"]

	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,               // 禁用自动创建外键约束
		Logger:                                   getGormLogger(cfg), // 使用自定义 Logger
	}); err != nil {
		zapLogger.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db, nil
	}
}
