package logx

type Config struct {
	ConsoleAppenderConfig *ConsoleAppenderConfig
	FileAppenderConfigs   []FileAppenderConfig
}

type CommonAppenderConfig struct {
	Level Level
}

type ConsoleAppenderConfig struct {
	CommonAppenderConfig
}

type FileAppenderConfig struct {
	CommonAppenderConfig
	FilePath string
}
