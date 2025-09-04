package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
)

type Client struct {
	IPv6       bool   `toml:"ipv6"`
	MixedPort  int    `toml:"mixed-port"`
	ManagePort int    `toml:"manage-port"`
	Secret     string `toml:"secret"`
	AllowLAN   bool   `toml:"allow-lan"`
	LogLevel   string `toml:"log-level"`
	Mode       string `toml:"mode"`
}
type Check struct {
	Enable   bool   `toml:"enable"`
	URL      string `toml:"url"`
	Interval int    `toml:"interval"`
	Timeout  int    `toml:"timeout"`
}

type config struct {
	Client Client `toml:"client"`
	Check  Check  `toml:"check"`
}

func main() {
	lines := []string{
		" __  __       _           ",
		"|  \\/  | ___ (_)_  ____ _  ",
		"| |\\/| |/ _ \\| \\ \\/ / _` |",
		"| |  | | (_) | |>  < (_| |",
		"|_|  |_|\\___/|_/_/\\_\\__,_|",
	}
	startColor := 19
	for i, line := range lines {
		colorCode := startColor + i*10
		fmt.Printf("\033[38;5;%dm%s\033[0m\n", colorCode, line)
	}
	color.Cyan("A high-performance , security network proxy tool.")
	configPath := flag.String("c", "./config.toml", "Path to config file")
	flag.Parse()

	var cfg config
	if _, err := toml.DecodeFile(*configPath, &cfg); err != nil {
		color.Red("Failed to load config file!")
		color.Red(err.Error())
		os.Exit(0)
	}

	color.Green("Config loaded successfully!")
	color.Green("Client part:")
	fmt.Printf("IPv6 Support: %v\n", cfg.Client.IPv6)
	fmt.Printf("Mixed Port: %d\n", cfg.Client.MixedPort)
	fmt.Printf("Manage Port: %d\n", cfg.Client.ManagePort)
	fmt.Printf("Secret: %s\n", cfg.Client.Secret)
	fmt.Printf("Allow LAN: %v\n", cfg.Client.AllowLAN)
	fmt.Printf("Log Level: %s\n", cfg.Client.LogLevel)
	fmt.Printf("Mode: %s\n", cfg.Client.Mode)
	color.Green("Connectivity check section:")
	fmt.Printf("Enable: %v\n", cfg.Check.Enable)
	fmt.Printf("URL: %s\n", cfg.Check.URL)
	fmt.Printf("Interval: %d\n", cfg.Check.Interval)
	fmt.Printf("Timeout: %d\n", cfg.Check.Timeout)
}
