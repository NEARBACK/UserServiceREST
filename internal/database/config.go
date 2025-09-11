package database

type PostgresConfig struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true" default:"postgres"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" required:"true" default:"5432"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true" default:"postgres"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true" default:"postgres"`
	PostgresDB       string `envconfig:"POSTGRES_DB" required:"true" default:"postgres"`
}

type config struct {
	PostgresMasterAddr string
	PostgresSlaveAddr  string
}

func (c *PostgresConfig) UrlConfig() config {
	return config{
		PostgresMasterAddr: "postgresql://" + c.PostgresUser + ":" + c.PostgresPassword + "@" + c.PostgresHost + ":" + c.PostgresPort + "/" + c.PostgresDB + "?sslmode=disable",
		PostgresSlaveAddr:  "postgresql://" + c.PostgresUser + ":" + c.PostgresPassword + "@" + c.PostgresHost + ":" + c.PostgresPort + "/" + c.PostgresDB + "?sslmode=disable",
	}
}
