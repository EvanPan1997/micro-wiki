package config

type Mysql struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Address  string `mapstructure:"address" yaml:"address"`
	Port     string `mapstructure:"port" yaml:"port"`
	DbName   string `mapstructure:"dbname" yaml:"dbname"`
	Config   string `mapstructure:"config" yaml:"config"`
}

//func (m *Mysql) Dsn() string {
//	return m.Username + ":" + m.Password + "@tcp(" + m.Address + ":" + string(rune(m.Port)) + ")/" + m.DbName + "?" + m.Config
//}
