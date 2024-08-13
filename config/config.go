package config

import "time"

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	ConnMaxLifetime time.Duration
}
