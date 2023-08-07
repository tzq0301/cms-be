package config

type DB struct {
	Master    DBInstance `mapstructure:"master"`
	Databases []Database `mapstructure:"databases"`
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
