package database

import (
	"context"
	"log"
	"useservice/internal/definitions"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type ContextKey string

var (
	dbConf = new(PostgresConfig)
)

// Init initialization from environment variables
func init() {
	if err := envconfig.Process("", dbConf); err != nil {
		log.Fatalf("db layer failed to load configuration: %s", err)
	}
}

var _ IDatabase = (*Database)(nil)

func New(log definitions.Logger) (*Database, error) {
	dbl := &Database{log: log}
	var err error
	conf := dbConf.UrlConfig()
	log.Info("Connect to master postgresql")
	masterConfig, err := pgxpool.ParseConfig(conf.PostgresMasterAddr)
	if err != nil {
		return nil, err
	}
	masterConfig.ConnConfig.PreferSimpleProtocol = true
	dbl.writePool, err = pgxpool.ConnectConfig(context.Background(), masterConfig)
	if err != nil {
		return nil, err
	}
	log.Info("Connected...")

	log.Info("Connect to slaves postgresql")
	slaveConfig, err := pgxpool.ParseConfig(conf.PostgresSlaveAddr)
	if err != nil {
		return nil, err
	}
	slaveConfig.ConnConfig.PreferSimpleProtocol = true
	dbl.readPool, err = pgxpool.ConnectConfig(context.Background(), slaveConfig)
	if err != nil {
		return nil, err
	}
	log.Info("Connected...")

	return dbl, nil
}
