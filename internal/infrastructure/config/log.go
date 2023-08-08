package config

type Log struct {
	Console *Console `mapstructure:"console"`
	File    []File   `mapstructure:"file"`
}

type Console struct {
	Level string
}

type File struct {
	Level    string
	FilePath string
}
