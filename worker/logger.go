package worker

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerRedisTask struct{}

// TODO: This logger is write file, it make u can check response from Redis.
func NewLoggerRedisTask() *LoggerRedisTask {
	return &LoggerRedisTask{}
}

func (logger *LoggerRedisTask) Print(level zerolog.Level, msg ...interface{}) {
	log.WithLevel(level).Msg(fmt.Sprint(msg...))
}

func (logger *LoggerRedisTask) Debug(args ...interface{}) {
	logger.Print(zerolog.DebugLevel, args...)
}

func (logger *LoggerRedisTask) Info(args ...interface{}) {
	logger.Print(zerolog.InfoLevel, args...)
}

func (logger *LoggerRedisTask) Warn(args ...interface{}) {
	logger.Print(zerolog.WarnLevel, args...)
}

func (logger *LoggerRedisTask) Error(args ...interface{}) {
	logger.Print(zerolog.ErrorLevel, args...)
}

func (logger *LoggerRedisTask) Fatal(args ...interface{}) {
	logger.Print(zerolog.FatalLevel, args...)
}
