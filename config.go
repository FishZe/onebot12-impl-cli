package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Onebot struct {
		Host              string `json:"host"`
		Port              int    `json:"port"`
		AccessToken       string `json:"access_token"`
		ReconnectInterval int    `json:"reconnect_interval"`
		UserAgent         string `json:"user_agent"`
		Impl              string `json:"impl"`
		Version           string `json:"version"`
	}
	Self struct {
		UserId   string `json:"user_id"`
		Platform string `json:"platform"`
		UserName string `json:"user_name"`
	}
}

func makeConfig() {
	f, err := os.Create("./config.yaml")
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("close file failed: %v", err)
		}
	}(f)
	if err != nil {
		log.Printf("create file failed: %v", err)
	} else {
		s, err := yaml.Marshal(&Config{})
		if err != nil {
			log.Printf("marshal config failed: %v", err)
			return
		}
		_, err = f.WriteString(string(s))
	}
}

func readConfig() (*Config, error) {
	f, err := os.Open("./config.yaml")
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("close file failed: %v", err)
		}
	}(f)
	if err != nil {
		log.Printf("open file failed: %v", err)
		return nil, err
	}
	var config Config
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		log.Printf("decode config failed: %v", err)
		return nil, err
	}
	return &config, nil
}
