package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	IP []string `yaml:"ip"`
}

func (c *config) getConfig() {
	file, err := ioutil.ReadFile("ip.yaml")
	if err != nil {
		log.Printf("file.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func main() {
	var  c config
	c.getConfig()
	fmt.Println(len(c.IP))
}
