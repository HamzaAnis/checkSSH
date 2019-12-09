package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	IP []string `yaml:"ip"`
}

func (c *config) getConfig(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Invalid config file   #%v ", err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func checkSSH(hosts []string, port string, successChan, errorChan chan string) {
	defer wg.Done()
	for _, host := range hosts {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			errorChan <- host
		}
		if conn != nil {
			successChan <- host
			conn.Close()
		}
	}
}

func perform(configFileName string) {
	var wg sync.WaitGroup
	successChan := make(chan string)
	errorChan := make(chan string)

	var c config
	c.getConfig(configFileName)

	threads := 10
	var splits [][]string

	chunk := (len(c.IP) + threads - 1) / threads

	for i := 0; i < len(c.IP); i += chunk {
		end := i + chunk

		if end > len(c.IP) {
			end = len(c.IP)
		}

		splits = append(splits, c.IP[i:end])
	}
	count := 0
	go func() {
		successFile, err := os.OpenFile("success.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		errorFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer successFile.Close()
		defer errorFile.Close()

		for i := 0; i < len(c.IP); i++ {
			select {
			case errorIP := <-errorChan:
				fmt.Println(errorIP + " [FAIL]")
				if _, err = errorFile.WriteString(errorIP + " [FAIL]\n"); err != nil {
					panic(err)
				}
			case successIP := <-successChan:
				if _, err = errorFile.WriteString(successIP + " [SUCCESS]\n"); err != nil {
					panic(err)
				}
				fmt.Println(successIP + " [SUCCESS")
			}
			count++
		}
	}()

	for _, split := range splits {
		go checkSSH(split, "22", successChan, errorChan)
		wg.Add(1)
	}
	wg.Wait()
}
