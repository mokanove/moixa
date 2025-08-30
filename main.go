package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Client struct {
	MixedPort  int    `toml:"mixed-port"`
	ManagePort int    `toml:"manage-port"`
	Secret     string `toml:"secret"`
	AllowLAN   bool   `toml:"allow-lan"`
	LogLevel   string `toml:"log-level"`
	Mode       string `toml:"mode"`
}

type config struct {
	Client Client `toml:"client"`
}

func main() {
	fmt.Println("Moixa:A high-performance network proxy tool.")
	configPath := flag.String("c", "./config.toml", "Path to config file")
	flag.Parse()

	var cfg config
	if _, err := toml.DecodeFile(*configPath, &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("Config loaded successfully:")
	fmt.Printf("Mixed Port: %d\n", cfg.Client.MixedPort)
	fmt.Printf("Manage Port: %d\n", cfg.Client.ManagePort)
	fmt.Printf("Secret: %s\n", cfg.Client.Secret)
	fmt.Printf("Allow LAN: %v\n", cfg.Client.AllowLAN)
	fmt.Printf("Log Level: %s\n", cfg.Client.LogLevel)
	fmt.Printf("Mode: %s\n", cfg.Client.Mode)

}
