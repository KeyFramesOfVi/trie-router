package env

import "github.com/caarlos0/env"

// setup for SQL server
const (
	DbDriver = "postgres"
)

// Config container
type Config struct {
	DevEnv bool   `env:"DEV_ENV" envDefault:"true"`
	Port   string `env:"SERVER_PORT" envDefault:"8080"`
	DB     DB
}

// DB conf
type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Username string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	Name     string `env:"POSTGRES_DB" envDefault:"fe_calc"`
}

// NewConfig function
func NewConfig() *Config {
	conf := &Config{}

	err := env.Parse(conf)
	if err != nil {
		panic(err)
	}

	err = env.Parse(&conf.DB)
	if err != nil {
		panic(err)
	}

	return conf
}

//TestConfig test server conf
type TestConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
	DB   TestDB
}

//TestDB conf
type TestDB struct {
	Host     string `env:"TEST_DB_HOST" envDefault:"postgres"`
	Username string `env:"TEST_POSTGRES_USER" envDefault:"wintermute"`
	Password string `env:"TEST_POSTGRES_PASSWORD" envDefault:"t0b30rn0tt0b3"`
	Name     string `env:"TEST_POSTGRES_DB" envDefault:"nypm_test"`
}

//NewTestConfig parse environment variables into Config
func NewTestConfig() *TestConfig {
	conf := &TestConfig{}

	// parse server conf
	err := env.Parse(conf)
	if err != nil {
		panic(err)
	}

	// parse DB conf
	err = env.Parse(&conf.DB)
	if err != nil {
		panic(err)
	}

	return conf
}
