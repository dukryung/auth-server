package configs

type LogConfig struct {
	Enable bool `json:"enable"`
	Level  int  `json:"level"`
}
