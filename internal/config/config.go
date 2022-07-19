package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	DbTypeSQL   = "SQL"
	DbTypeNoSql = "NoSQL"
)

// envVars contains all supported env with their default values.
// It's needed for viper to properly read environmental variables into config.
var envVars = []struct {
	name         string
	defaultValue string
}{
	{"PORT", "8080"},
	{"DB_TYPE", "SQL"},
	{"PSQL_DB_PORT", ""},
	{"PSQL_DB_USER", ""},
	{"PSQL_DB_PASSWORD", ""},
	{"PSQL_DB_HOST", ""},
	{"PSQL_DB_NAME", ""},
	{"MONGO_PORT", ""},
	{"MONGO_DB_HOST", ""},
	{"MONGO_USERNAME", ""},
	{"MONGO_PASSWORD", ""},
	{"MONGO_DB_NAME", ""},
}

// Config holds whole app configuration.
type Config struct {
	Port   int    `mapstructure:"PORT"`
	DbType string `mapstructure:"DB_TYPE"` // SQL or NoSQL.

	PsqlDbPort     int    `mapstructure:"PSQL_DB_PORT"`
	PsqlDbUser     string `mapstructure:"PSQL_DB_USER"`
	PsqlDbPassword string `mapstructure:"PSQL_DB_PASSWORD"`
	PsqlDbHost     string `mapstructure:"PSQL_DB_HOST"`
	PsqlDbName     string `mapstructure:"PSQL_DB_NAME"`

	MongoPort     int    `mapstructure:"MONGO_PORT"`
	MongoDbHost   string `mapstructure:"MONGO_DB_HOST"`
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoDbName   string `mapstructure:"MONGO_DB_NAME"`
}

// PsqlConnString returns postgres connection string.
func (c Config) PsqlConnString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PsqlDbHost,
		c.PsqlDbPort,
		c.PsqlDbUser,
		c.PsqlDbPassword,
		c.PsqlDbName)
}

// MongoConnString returns mongo connection string.
func (c Config) MongoConnString() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		c.MongoUsername,
		c.MongoPassword,
		c.MongoDbHost,
		c.MongoPort,
	)
}

// InitConfig binds envs, reads them into structure and returns.
func InitConfig(configPath string) (*Config, error) {
	for _, env := range envVars {
		if err := viper.BindEnv(env.name); err != nil {
			return nil, err
		}
		if env.defaultValue != "" {
			viper.SetDefault(env.name, env.defaultValue)
		}
	}

	viper.SetEnvPrefix("")

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config, %w", err)
	}

	return &c, nil
}
