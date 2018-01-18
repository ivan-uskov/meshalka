package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Dsn    string `yaml:"dsn"`
	Driver string `yaml:"driver"`
}

type Config struct {
	Database      DBConfig `yaml:"database"`
	SessionSecret string   `yaml:"session_secret"`
	AppAddress    string   `yaml:"listen"`
}

func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal([]byte(data), c)
	return c, err
}