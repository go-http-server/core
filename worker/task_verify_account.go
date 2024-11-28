package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5"
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

	log.Printf("ENQUEUED_TASK: [TYPE]: %s\n[PAYLOAD]: %s\n[QUEUE]: %s\n[MAX_RETRY]: %d", task.Type(), task.Payload(), taskInfo.Queue, taskInfo.MaxRetry)

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendEmailVerifyAccount(ctx context.Context, task *asynq.Task) error {
	var payload mailer.UserReceive
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("[ERROR_PROCESS_TASK] - [%s]: failed Unmarshal payload: %s %w", TaskSendVerifyAccount, err, asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, database.GetUserParams{
		Username: payload.Username,
		Email:    payload.EmailAddress,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("[ERROR_PROCESS_TASK] - [%s]: Not found user %w", TaskSendVerifyAccount, asynq.SkipRetry)
		}

		return fmt.Errorf("[ERROR_PROCESS_TASK] - [%s]: Failed get user: %s", TaskSendVerifyAccount, err)
	}

	// TODO: Send email to user
	log.Printf("PROCESS_TASK: [TYPE]: %s\n[PAYLOAD]: %s\n[EMAIL]: %s", task.Type(), task.Payload(), user.Email)

	return nil
}
