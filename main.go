package main

import (
	"MEND/internal/config"
	"MEND/pkg/repository"
	"MEND/pkg/server"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const configPathEnv = "CONFIG_PATH"

func main() {
	os.Exit(run())
}

func run() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := log.Default() // NOTE: In production I would use something more advanced like zerolog or logrus.

	configPath := os.Getenv(configPathEnv)

	config, err := config.InitConfig(configPath)
	if err != nil {
		log.Printf("failed to init config, %s\n", err)
		return 1
	}

	log.Println(config)

	repo, err := createRepository(ctx, config)
	if err != nil {
		log.Printf("failed to init repository, %s\n", err)
		return 1
	}
	controller := server.NewUserController(repo)

	router := server.NewRouter(controller)

	if err := router.RunTLS(fmt.Sprintf(":%d", config.Port), "server.crt", "server.key"); err != nil {
		log.Printf("error during router run %s\n", err)
		return 1
	}
	return 0
}

// createRepository takes config and based on DbType creates UserRepository.
func createRepository(ctx context.Context, cfg *config.Config) (server.UserRepository, error) {
	var repo server.UserRepository

	switch cfg.DbType {
	case config.DbTypeSQL:
		db, err := sql.Open("postgres", cfg.PsqlConnString())
		if err != nil {
			return nil, fmt.Errorf("failed to open sql connection, %w", err)
		}
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to ping database, %w", err)
		}
		repo = repository.NewPSQL(db)
	case config.DbTypeNoSql:
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoConnString()))
		if err != nil {
			return nil, fmt.Errorf("failed to connect to mongo db, %w", err)
		}

		// Ping the primary
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("failed to ping mongo database, %w", err)
		}

		repo = repository.NewMongo(client.Database(cfg.MongoDbName).Collection("user"))

	default:
		return nil, fmt.Errorf("unknown db type: %s", cfg.DbType)
	}

	return repo, nil
}
