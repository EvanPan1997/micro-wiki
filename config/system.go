package config

type Server struct {
	DbType string `mapstructure:"db-type" yaml:"db-type"`
}
