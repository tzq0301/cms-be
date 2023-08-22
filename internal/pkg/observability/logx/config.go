package logx

type Config struct {
	ServiceConfig         ServiceConfig
	ConsoleAppenderConfig *ConsoleAppenderConfig
	FileAppenderConfigs   []FileAppenderConfig
}

type ServiceConfig struct {
	Name string   `json:"name"`
	IP   IPConfig `json:"IP,omitempty"`
}

type IPConfig struct {
	V4 string `json:"v4,omitempty"`
	V6 string `json:"v6,omitempty"`
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
