package config

type Adaptor struct {
	Gin Gin `mapstructure:"gin"`
}

type Gin struct {
	Port      int    `mapstructure:"port"`
	UrlPrefix string `mapstructure:"urlPrefix"`
}
