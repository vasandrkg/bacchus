// some comment
package main

import (
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
}

func (c *Config) readConfig() *Config {

	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshaling error: %v", err)
	}

	return c
}

func main() {
	var c Config
	c.readConfig()

	fmt.Println(c)
}
