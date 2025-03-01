package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName                string        `mapstructure:"APP_NAME"`
	AppVersion             string        `mapstructure:"APP_VERSION"`
	AppHTTPPort            int           `mapstructure:"APP_HTTP_PORT"`
	AppGRPCPort            int           `mapstructure:"APP_GRPC_PORT"`
	DBDriver               string        `mapstructure:"DB_DRIVER"`
	DBUsername             string        `mapstructure:"DB_USERNAME"`
	DBPassword             string        `mapstructure:"DB_PASSWORD"`
	DBName                 string        `mapstructure:"DB_NAME"`
	DBHost                 string        `mapstructure:"DB_HOST"`
	DBPort                 int           `mapstructure:"DB_PORT"`
	DBMaxIdleConn          int           `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn          int           `mapstructure:"DB_MAX_OPEN_CONN"`
	DBMaxConnLifetime      time.Duration `mapstructure:"DB_MAX_CONN_LIFETIME"`
	DBSslMode              string        `mapstructure:"DB_SSL_MODE"`
	DBEnableDebug          bool          `mapstructure:"DB_ENABLE_DEBUG"`
	LoggerEnable           bool          `mapstructure:"LOGGER_ENABLE"`
	LoggerEnableStackTrace bool          `mapstructure:"LOGGER_ENABLE_STACK_TRACE"`
	LoggerEnableMasking    bool          `mapstructure:"LOGGER_ENABLE_MASKING"`
	LoggerMaskingFields    []string      `mapstructure:"LOGGER_MASKING_FIELDS"`
	LocationSvcUrl         string        `mapstructure:"LOCATION_SVC_URL"`
	LocationSvcTimeout     time.Duration `mapstructure:"LOCATION_SVC_TIMEOUT"`
	LocationSvcSkipTLS     bool          `mapstructure:"LOCATION_SVC_SKIP_TLS"`
	RabbitMQHost           string        `mapstructure:"RABBIT_MQ_HOST"`
	RabbitMQPort           int           `mapstructure:"RABBIT_MQ_PORT"`
	RabbitMQUsername       string        `mapstructure:"RABBIT_MQ_USERNAME"`
	RabbitMQPassword       string        `mapstructure:"RABBIT_MQ_PASSWORD"`
}

func Load() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return conf
}

func (c *Config) GetPostgresDSN() (dsn string) {
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.DBHost,
		c.DBUsername,
		c.DBPassword,
		c.DBName,
		c.DBPort,
		c.DBSslMode,
	)

	return
}
