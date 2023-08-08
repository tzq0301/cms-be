package config

type MySQL struct {
	Name     string        `mapstructure:"name"`
	Master   MySQLInstance `mapstructure:"master"`
	Database []Database    `mapstructure:"database"`
}

type MySQLInstance struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Database struct {
	Name string
}
