package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server Server `yaml:"server"`
}

// Server
type Server struct {
	Timeout Timeout `yaml:"timeout"`
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
}

// Timeout
type Timeout struct {
	Server int `yaml:"server"`
	Read   int `yaml:"read"`
	Write  int `yaml:"write"`
	Idle   int `yaml:"idle"`
}

func GetAppConfig() (*AppConfig, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(filepath.Join(wd, "pkg", "config", "app.yml"))
	if err != nil {
		return nil, err
	}
	app := &AppConfig{}
	err = yaml.Unmarshal(b, app)
	if err != nil {
		return nil, err
	}

	return app, err
}
