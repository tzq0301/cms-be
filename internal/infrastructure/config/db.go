package config

type DB struct {
	Name     string     `mapstructure:"name"`
	Master   DBInstance `mapstructure:"master"`
	Database []Database `mapstructure:"database"`
}

type DBInstance struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Database struct {
	Name string
}
