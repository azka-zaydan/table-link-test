package configs

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		CORS struct {
			AllowCredentials bool   `mapstructure:"ALLOW_CREDENTIALS"`
			AllowedHeaders   string `mapstructure:"ALLOWED_HEADERS"`
			AllowedMethods   string `mapstructure:"ALLOWED_METHODS"`
			AllowedOrigins   string `mapstructure:"ALLOWED_ORIGINS"`
			Enable           bool   `mapstructure:"ENABLE"`
			MaxAgeSeconds    int    `mapstructure:"MAX_AGE_SECONDS"`
		}
		Name     string `mapstructure:"NAME"`
		Revision string `mapstructure:"REVISION"`
		URL      string `mapstructure:"URL"`
	}

	Cache struct {
		Redis struct {
			Primary struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Password string `mapstructure:"PASSWORD"`
				DB       int    `mapstructure:"DB"`
			}
		}
		Token struct {
			ExpiresIn time.Duration `mapstructure:"EXPIRES_IN"`
		} `mapstructure:"TOKEN"`
		Cart struct {
			ExpiresIn time.Duration `mapstructure:"EXPIRES_IN"`
		} `mapstructure:"CART"`
	}

	DB struct {
		MySQL struct {
			Read struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
			Write struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
		}

		Postgre struct {
			Read struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
			Write struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
		}
	}

	JWT struct {
		ExpiresIn time.Duration `mapstructure:"EXPIRES_IN"`
		Key       string        `mapstructure:"KEY"`
	} `mapstructure:"JWT"`

	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Port     string `mapstructure:"PORT"`
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		}
		GRPC struct {
			Host    string `mapstructureL:"HOST"`
			Network string `mapstructure:"NETWORK"`
			Port    string `mapstructure:"PORT"`
		} `mapstructure:"GRPC"`
	}
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	return &conf
}
