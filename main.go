package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ParamInfo struct {
	Parameter_name  string `yaml:"PARAMETER_NAME"`
	Is_default      string `yaml:"IS_DEFAULT"`
	Parameter_type  string `yaml:"PARAMETER_TYPE"`
	Parameter_value string `yaml:"PARAMETER_VALUE"`
}
type ServerInfo struct {
	Server              string      `yaml:"SERVER"`
	Server_display_name string      `yaml:"SERVER_DISPLAY_NAME"`
	Server_port         int64       `yaml:"SERVER_PORT"`
	Server_params       []ParamInfo `yaml:"SERVER_PARAMS"`
}

func main() {
	yfile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data []ServerInfo
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Print(data)

}
