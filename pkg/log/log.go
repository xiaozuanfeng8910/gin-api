package log

import (
	"gin-api/internal/config"
	"gin-api/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func createRootDir(rootDir string) {
	if ok := utils.PathExists(rootDir); !ok {
		_ = os.Mkdir(rootDir, os.ModePerm)
	}
}

func setLogLevel(_level string) {
	switch _level {
	case "debug":
		level := zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore(log *config.LogConfig) zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(l.String())
	}

	// 设置编码器
	if log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(log), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter(log *config.LogConfig) zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   log.RootDir + "/" + log.Filename,
		MaxSize:    log.MaxSize,
		MaxBackups: log.MaxBackups,
		MaxAge:     log.MaxAge,
		Compress:   log.Compress,
	}

	return zapcore.AddSync(file)
}

func InitializerLog(cfg *config.Config) (*zap.Logger, error) {
	// 创建根目录
	createRootDir(cfg.Log.RootDir)

	// 设置日志等级
	setLogLevel(cfg.Log.Level)

	if cfg.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(&cfg.Log), options...), nil
}
