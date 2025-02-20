package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	AuthSecret string `json:"auth_secret"`
	Listen     string `json:"listen"`
	HookUri    string `json:"hook_uri"`
	UpdateUri  string `json:"update_uri"`
	WebUri     string `json:"web_uri"`
	HookToken  string `json:"hook_token"`
}

var cfg *Config

func LoadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Panic(err)
	}

	cfg = &Config{}
	err = json.Unmarshal(file, cfg)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(">>>加载配置：", cfg)

	return
}
