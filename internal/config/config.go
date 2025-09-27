package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config read only.
type Config struct {
	Debug                 bool          `envconfig:"DEBUG" default:"true"`
	ServicePort           string        `envconfig:"PORT" default:"8000"`
	HealthCheckPort       int           `envconfig:"HEALTHCHECK_PORT" default:"8001"`
	SecretAccessTokenJWT  string        `envconfig:"SECRET_ACCESS_TOKEN_JWT" default:""`
	SecretRefreshTokenJWT string        `envconfig:"SECRET_REFRESH_TOKEN_JWT" default:""`
	AccessTtlMinuteJWT    time.Duration `envconfig:"ACCESS_TTL_MINUTE_JWT" default:"15"`
	RefreshTtlDayJWT      time.Duration `envconfig:"REFRESH_TTL_DAY_JWT" default:"7"`
}

// New Config constructor.
func New() *Config {
	return &Config{}
}

// Init initialization from environment variables
func (r *Config) Init() {
	if err := envconfig.Process("", r); err != nil {
		log.Printf("failed to load configuration: %s", err)
	}
}
