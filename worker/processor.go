package worker

import (
	"context"
	"fmt"

	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/go-http-server/core/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	Shutdown()
	ProcessTaskSendEmailVerifyAccount(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  database.Store
	sender mailer.EmailSender
}

func NewRedisTaskProcessor(redisOpts asynq.RedisClientOpt, store database.Store, sender mailer.EmailSender, bot *utils.BotTelegram) TaskProcessor {
	server := asynq.NewServer(redisOpts, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Str("[PROCESS_TASK]", TaskSendVerifyAccount).
				Str("[TYPE]", task.Type()).
				Bytes("[PAYLOAD]", task.Payload())

			messageToBot := fmt.Sprintf("[PROCESS_TASK] - [%s]: [TYPE]: %s\n[PAYLOAD]: %s", TaskSendVerifyAccount, task.Type(), task.Payload())
			bot.SendMessage(messageToBot)
		}),
		Logger: NewLoggerRedisTask(),
	})

	return &RedisTaskProcessor{
		server: server,
		store:  store,
		sender: sender,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	serverMux := asynq.NewServeMux()
	serverMux.HandleFunc(TaskSendVerifyAccount, processor.ProcessTaskSendEmailVerifyAccount)

	return processor.server.Start(serverMux)
}

func (processor *RedisTaskProcessor) Shutdown() {
	processor.server.Shutdown()
}
