package logx

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var sugarLogger *zap.SugaredLogger

type LogxLvl string

func Setup(logPath string, level LogxLvl) {
	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.ErrorLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    10,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		}),
		logLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	sugarLogger = logger.Sugar()

	log.Println("\t\t" + color.GreenString("[OK]: ") + "logging setup successfully.")
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", level.CapitalString()))
}

func Sync() error {
	return sugarLogger.Sync()
}

func Debug(args ...any) {
	sugarLogger.Debugln(args...)
}

func Info(args ...any) {
	sugarLogger.Infoln(args...)
}

func Warn(args ...any) {
	sugarLogger.Warnln(args...)
}

func Error(args ...any) {
	sugarLogger.Errorln(args...)
}

func Fatal(args ...any) {
	sugarLogger.Fatalln(args...)
}
