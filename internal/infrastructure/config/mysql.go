package config

type MySQL struct {
	Database string          `mapstructure:"database"`
	Master   MySQLInstance   `mapstructure:"master"`
	Slaves   []MySQLInstance `mapstructure:"slaves"`
}

type MySQLInstance struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
