package worker

import (
	"context"

	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendEmailVerifyAccount(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  database.Store
}

func NewRedisTaskProcessor(redisOpts asynq.RedisClientOpt, store database.Store) TaskProcessor {
	server := asynq.NewServer(redisOpts, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
	})

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	serverMux := asynq.NewServeMux()
	serverMux.HandleFunc(TaskSendVerifyAccount, processor.ProcessTaskSendEmailVerifyAccount)

	return processor.server.Start(serverMux)
}
