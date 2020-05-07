package config

type LoggerConfig struct {
	Output string `json:"output"`
}

type Config struct {
	TemplatePath string `json:"template_path"`
	Logger LoggerConfig `json:"logger"`
}

func NewConfig(templatePath string) *Config {
	return &Config{TemplatePath: templatePath}
}
