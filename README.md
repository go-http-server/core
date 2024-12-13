### INTRODUCE

- This is a simple backend service with the purpose of learning and building a foundation for other backend projects.
- In here, I rewrite some fundamental function like register, login and implements auth middleware.

### FEATURES

- Pass all test cases.
- With authentication:
  - Register:
    - Create user, have validation with `validator/v10`.
    - Send code verify to mail with flow: `endpoint register` -> push to `redis queue` and return response early to user -> send mail with `code` (use template HTML).
    - Implements `transaction in database`, it make you can `rollback` or `commit` transaction.
  - Login:
    - Sign token: I use `PASETO` to generate token with `asymmetricKey`. It use `privateKey` to sign token, use `publicKey` to verify token. `JWT` is **buffoonery**
  - Middleware: Only support `Bearer` token and apply to necessary route.
- Building: you can use Docker or build with param `mod vendor`. Folder `vendor` like `node_modules` in Nodejs, it is folder contains packages used -> avoid missing or conflict version.

### TOOLS

- `sqlc`: generate Go performs T-SQL.
- `redoc-cli`: to generate `yml` docs file.

### ENVIRONMENTS

- **Go**: version 1.23.0
- **Docker**: I like to use docker-compose
- **sqlc**: version 1.27.0 or above
- **Database**: Postgresql 17

- Variables environments:
  - `DB_SOURCE`: connection string to Postgresql database like: `postgresql://postgres:phamnam123@localhost:5432/core?sslmode=disable`
  - `HTTP_SERVER_ADDRESS`: Server Address running like: `0.0.0.0:8080`
  - `ENVIRONMENT`: Mode application running like: `develop` or `product`
  - `TIME_EXPIRED_TOKEN`: Time token will be expired like: `15m`
  - `EMAIL_PASSWORD_SENDER`: Google app password (not real password your email) like: `cgqtvyhydtmdihep`
  - `EMAIL_ADDRESS_SENDER`: Email sender like: `namphamhai7@gmail.com`
  - `EMAIL_USERNAME_SENDER`: Username makes your email look like professional like: `Go Core Service`
  - `REDIS_ADDRESS_SERVER`: Address of Redis, like: `0.0.0.0:6379`
  - `REDIS_PASSWORD_SERVER`: Password for access to Redis database, like: `redis_password`. If you use my `docker-compose`, this var must be equal var `REDIS_PASSWORD` in file `docker-compose.yml`
  - `TELEGRAM_BOT_TOKEN`: Token of your bot in Telegram, look like: `7733600827:AAE_s4H5M1Vh9n4BKfz4YgImf-q9ltMhDSA`
  - `TELEGRAM_CHAT_ID`: Your chat id, like: `6197489552`

### Usage

- Install `sqlc`:
- With **Arch Linux**

```bash
sudo pacman -S sqlc
```

- With **Ubuntu or Debian**

```bash
sudo apt install sqlc
```

- With **MACOS**

```bash
brew install sqlc
```

- Run project:

```bash
go run cmd/server.go -mod`=vendor
```
