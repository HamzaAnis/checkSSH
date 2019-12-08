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

	threads:=10
	var split [][]string

	chunk := (len(c.IP) + threads - 1) / threads

	for i := 0; i < len(c.IP); i += chunk {
		end := i + chunk

		if end > len(c.IP) {
			end = len(c.IP)
		}

		split = append(split, c.IP [i:end])
	}

	fmt.Printf("%#v\n", split)
	fmt.Printf("%#v\n", len(split))
	fmt.Println(len(c.IP))
}
