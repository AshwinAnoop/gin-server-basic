package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type ParamInfo struct {
	Parameter_name  string `yaml:"PARAMETER_NAME"`
	Is_default      string `yaml:"IS_DEFAULT"`
	Parameter_type  string `yaml:"PARAMETER_TYPE"`
	Parameter_value string `yaml:"PARAMETER_VALUE"`
	Level           string
	Rank            int64
}
type ServerInfo struct {
	Server              string      `yaml:"SERVER"`
	Server_display_name string      `yaml:"SERVER_DISPLAY_NAME"`
	Server_port         int64       `yaml:"SERVER_PORT"`
	Server_params       []ParamInfo `yaml:"SERVER_PARAMS"`
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/index/server", getServerDetails)
	router.GET("/index", mainPage)
	router.Run(":8080")
}

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

// Get handler to return comined data of yaml and json file
func getServerDetails(c *gin.Context) {

	yfile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data []ServerInfo
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	jfile, err3 := ioutil.ReadFile("test.json")
	if err3 != nil {
		log.Fatal(err3)
	}

	var jdata map[string]interface{}
	err4 := json.Unmarshal(jfile, &jdata)
	if err4 != nil {
		log.Fatal(err4)
	}

	fmt.Print("jdata : ", jdata)
	var paramName string
	for i := range data {
		for j := range data[i].Server_params {
			paramName = data[i].Server_params[j].Parameter_name
			if val, ok := jdata[paramName]; ok {

				m, _ := val.(map[string]interface{})

				fmt.Print(m)
				for k, v := range m {
					if k == "Level" {
						data[i].Server_params[j].Level = fmt.Sprint(v)

					}
					if k == "Rank" {
						data[i].Server_params[j].Rank = int64(v.(float64))
					}
				}
			}
		}
	}

	server := c.Query("server")
	for _, a := range data {
		if a.Server == server {
			c.HTML(http.StatusOK, "serverInfo.tmpl", gin.H{
				"a": a,
			})

		}
	}

}
