package database

type config struct {
	PostgresMasterAddr string `envconfig:"POSTGRES_MASTER_ADDR" required:"true" default:"postgres://postgres:postgres@postgre:5432/postgres?standard_conforming_strings=on&sslmode=disable"`
	PostgresSlaveAddr  string `envconfig:"POSTGRES_SLAVE_ADDR" required:"false" default:"postgres://postgres:postgres@postgre:5432/postgres?standard_conforming_strings=on&sslmode=disable"`
}
