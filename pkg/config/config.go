package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var wd_conf string

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

func init() {
	var err error
	wd_conf, err = os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	wd_conf = filepath.Join(wd_conf, "pkg", "config")
}

func GetAppConfig() (*AppConfig, error) {
	var err error
	b, err := ioutil.ReadFile(filepath.Join(wd_conf, "app.yml"))
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

func GetDBConfig() {

}
