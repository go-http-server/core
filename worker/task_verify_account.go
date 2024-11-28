package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyAccount(
	ctx context.Context,
	payload *mailer.UserReceive,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("[ERROR_TASK] - [%s]: failed to marshal payload: %s", TaskSendVerifyAccount, err)
	}

	task := asynq.NewTask(TaskSendVerifyAccount, jsonPayload, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		return fmt.Errorf("[ERROR_ENQUEUE_CONTEXT] - [%s]: %s", TaskSendVerifyAccount, err)
	}

	log.Info().Str("[TYPE]", task.Type()).Str("[TASK]", TaskSendVerifyAccount).
		Bytes("[PAYLOAD]", task.Payload()).Str("[QUEUE]", taskInfo.Queue).
		Int("[MAX_RETRY]", taskInfo.MaxRetry).Msg("ENQUEUED_TASK")

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendEmailVerifyAccount(ctx context.Context, task *asynq.Task) error {
	var payload mailer.UserReceive
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("[ERROR_PROCESS_TASK] - [%s]: failed Unmarshal payload: %s %w", TaskSendVerifyAccount, err, asynq.SkipRetry)
	}

	log.Info().Str("[TYPE]", task.Type()).Str("[TASK]", TaskSendVerifyAccount).Bytes("[PAYLOAD]", task.Payload())

	err = processor.sender.SendWithTemplate(
		"[Go Core] Kích hoạt tài khoản",
		"./templates/verify_account.html",
		payload,
	)
	if err != nil {
		return fmt.Errorf("[ERROR_PROCESS_TASK] - [%s]: %s", TaskSendVerifyAccount, err)
	}

	return nil
}
