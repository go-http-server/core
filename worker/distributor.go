package worker

import (
	"context"

	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/hibiken/asynq"
)

const (
	TaskSendVerifyAccount = "task:send_mail_verify_account"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyAccount(
		ctx context.Context,
		payload *mailer.UserReceive,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpts)
	return &RedisTaskDistributor{
		client: client,
	}
}
