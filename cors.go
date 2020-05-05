// Package ezcors handles CORS configuration file rs.cors Options.
package ezcors

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CORS defines the supported properties.
type CORS struct {
	AllowedOrigins     []string `yaml:"allowedOrigins"`
	AllowCredentials   bool     `yaml:"allowCredentials"`
	AllowedMethods     []string `yaml:"allowedMethods"`
	AllowedHeaders     []string `yaml:"allowedHeaders"`
	ExposedHeaders     []string `yaml:"exposedHeaders"`
	MaxAge             int      `yaml:"maxAge"`
	OptionsPassthrough bool     `yaml:"optionsPassthrough"`
	Debug              bool     `yaml:"debug"`
}

// Config defines CORS configuration for every environment.
type Config map[string]CORS

// NewConfig decodes config file and returns CORS Config.
// The function will look for cors.yml file from the current directory. If
// nothing can found it will try to look into the config directory for possible
// cors.yml file.
func NewConfig() (Config, error) {
	file, err := os.Open("cors.yml")
	if err != nil {
		file, err := os.Open("config/cors.yml")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	defer file.Close()

	config := Config{}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
