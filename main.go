package main

import (
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

	server := c.Query("server")
	fmt.Print(server)
	for _, a := range data {
		if a.Server == server {
			fmt.Print(a)
			c.HTML(http.StatusOK, "serverInfo.tmpl", gin.H{
				"a": a,
			})

		}
	}

}
