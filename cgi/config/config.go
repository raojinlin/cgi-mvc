package config

type Config struct {
	TemplatePath string `json:"template_path"`
}

func NewConfig(templatePath string) *Config {
	return &Config{TemplatePath: templatePath}
}
