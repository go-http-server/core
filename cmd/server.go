package main

import (
	"context"
	"os"

	"github.com/go-http-server/core/api"
	database "github.com/go-http-server/core/internal/database/sqlc"
	"github.com/go-http-server/core/plugin/pkg/mailer"
	"github.com/go-http-server/core/utils"
	"github.com/go-http-server/core/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	env, err := utils.LoadEnviromentVariables("./")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load environment variables")
	}
	if env.ENVIRONMENT != utils.ProductionEnvironment {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	pool, err := pgxpool.New(context.Background(), env.DB_SOURCE)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create pool to database")
	}
	store := database.NewStore(pool)

	redisOpts := asynq.RedisClientOpt{
		Addr: env.REDIS_ADDRESS_SEVRER,
	}

	emailSender := mailer.NewGmailSender(env.EMAIL_USERNAME_SENDER, env.EMAIL_ADDRESS_SENDER, env.EMAIL_PASSWORD_SENDER)
	taskDistributor := worker.NewRedisTaskDistributor(redisOpts)

	bot := utils.NewBotTelegramService(env.TELEGRAM_BOT_TOKEN, env.TELEGRAM_CHAT_ID)

	go runTaskProcessor(redisOpts, store, emailSender, bot)

	server, err := api.NewServer(store, env, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create new server")
	}
	server.StartServer(env.HTTP_SERVER_ADDRESS)
}

func runTaskProcessor(redisOpts asynq.RedisClientOpt, store database.Store, sender mailer.EmailSender, bot *utils.BotTelegram) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpts, store, sender, bot)
	log.Info().Msg("Start task processor")

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed start task processor")
	}
}
