package config

type MySQL struct {
	Name     string        `mapstructure:"name"`
	Database string        `mapstructure:"database"`
	Master   MySQLInstance `mapstructure:"master"`
}

type MySQLInstance struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
