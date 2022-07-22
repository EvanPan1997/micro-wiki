package config

type System struct {
	Server Server `mapstructure:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" yaml:"mysql"`
}
