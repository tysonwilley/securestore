package config

import (
	"os"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	"fmt"
)

type parametersBag struct{
	Server struct{
		Host string `yaml:host`
		Port string `yaml:port`
	}
	Database struct{
		User     string `yaml:user`
		Password string `yaml:password`
		Host     string `yaml:host`
		Port     string `yaml:port`
		Database string `yaml:database`
	}
}

var configPath string = fmt.Sprintf("%s/Work/go/src/secureStore/config.yml", os.Getenv("HOME"))
var Parameters parametersBag

func init() {
	dat, err := ioutil.ReadFile(configPath)
	errorPanic(err)

	err = yaml.Unmarshal(dat, &Parameters)
	errorPanic(err)
}


func errorPanic(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
