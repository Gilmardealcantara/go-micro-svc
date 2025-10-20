# **GO MICROSERVICE TEMPLATE ʕ◔ϖ◔ʔ**

Getting started with your very own Go microservice template.

## How to use
- Create a project base on this template using tech deck
- Replace go-service-template to your service name

## **Sample Project Structure** 🌳

```bash
.
├── README.md
├── api
│   ├── health.go
│   ├── health_test.go
│   ├── hello.go
│   ├── hello_test.go
│   └── setuprouter.go
├── cmd
│   └── helloworld
│       ├── README.md
│       └── main.go
├── config
│   ├── README.md
│   └── config.go
├── db
│   ├── conn.go
│   └── conn_test.go
├── go.mod
├── go.sum
└── pkg
    └── log
        ├── log.go
        └── log_test.go
```

## **Peripherals / Maintenance** 👩‍💻🧑‍💻

- Database Libs
  - [golang-migrate](https://github.com/golang-migrate/migrate)
  - [gorm](https://gorm.io/)
- Message Queues Libs
  - [asynq](https://github.com/hibiken/asynq)
  - others?
- Docker
  - [docker](https://www.docker.com/products/docker-desktop)
- Testing
  - [Interfaces & Composition for Testing](https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/)
  - [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/)


## Locally
```shell
asdf plugin add golang
asdf plugin add swag
asdf plugin add golangci-lint
asdf plugin add pre-commit
asdf plugin add gomigrate
asdf install


./scripts/setup
make run

# For start db
docker compose up -d
make db_migrate_up
```
## **Development Environment**

Docker compose makes is easy to set up local dev environment with all necessary services and binaries.
Also to maintain parity with production. All you need is to have Docker installed on your machine.
The rest will be taken care of by Docker Compose.
If you choose not to use it for your application feel free to remove it.


### Build the app:

`docker-compose build`

If you see `403 Forbidden` error when running this command,
then you need do login into AWS ECR Default profile.

Run: `aws ecr get-login-password --region us-east-1 --profile default | docker login --username AWS --password-stdin 315120000506.dkr.ecr.us-east-1.amazonaws.com`

**NOTE:** You can check your profiles info by running `cat ~/.aws/credentials`

### Run the app:

`docker-compose up`

### Verify that app is running:

`localhost:8000/__healthcheck__`

## **Deployment** 🚀🚀🚀

Uncommenting the following files should technically deploy this sample application:

- `cmd/helloworld/.infra/staging.yaml`
- `.github/workflows/staging.yaml`

Ping #dev-microservice channel with any questions

You can then view the service logs in New Relic:
https://rpm.newrelic.com/accounts/4102727/applications/1038040774

## **Resources**

### Go

- [Zen of Go](https://dave.cheney.net/2020/02/23/the-zen-of-go)
- [Go PR Review Tips](https://github.com/golang/go/wiki/CodeReviewComments#variable-names)
- [Go Course - Intermediate](https://www.youtube.com/watch?v=iDQAZEJK8lI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
- [Avoid Global Vars](https://blog.canopas.com/approach-to-avoid-accessing-variables-globally-in-golang-2019b234762)

### Observablity

- [Intro to Tracing - Go specific](https://www.youtube.com/watch?v=idDu_jXqf4E)

## Swagger
```
asdf plugin add swag
asdf install
make doc
```

# Docker build dev
```sh
docker build . -t go-microsvc-template --build-arg GH_TOKEN=${GH_PACKAGE_TOKEN}
```

## integration tests with test container
```shell
# Unit tests
make test_u

# Integration tests with test container
make test_i
# for rancher users
make test_ir
```

## Db and migrate
```shell
# create db on port 5433
docker compose up

# Connect to docker Database
make db_conn

# Run all migrations
make db_migrate_up

# Rollback all migrations
make db_migrate_down

# Run migration command
make db_migrate_goto goto 3
make db_migrate_force force 3 # when problems happen

# Create new migration
make db_migrate_create name=new_migration_name

```

### Choices:
- [x] use standard hpp server mux for http server
- [x] use pgx5 for postgres db conn
- [x] Add test container setup for integration tests
- [x] **Structured logs using slog and NR**
- [x] Setup db migrations
- [x] Update pre-commit hooks
- [x] Add asdf dependencies
- [x] Update Quality pipelines (Add lint, unit and integration tests)
- [x] simplified PRD Dockerfile

### TODO
- [x] Setup db migrations
- [x] Fix pipelines
- [x] update health check endpoint with db ping and hostname result
- [x] Fix dockerfile and docker compose...
- [x] Review and improve swagger config
- [x] Add gracefully shutdown config
- [x] Add rest client with telemetry
- [ ] Add http base utils dor telemetry
- [ ] Setup/Test dockerfile dev
- [ ] Add open telemetry setup
- [ ] Add temporal setup
- 
