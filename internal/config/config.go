package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	config_path_flag = "config-path"
	config_path_env  = "CONFIG_PATH"
)

var (
	path = flag.String(config_path_flag, "", "path to configure file")
)

type Config struct {
	Env       string `yaml:"env" env-default:"local"`
	BotToken  string `yaml:"bot_token" env-required:"true"`
	CachePort string `yaml:"redis_port" env-required:"true"`
}

func LoadConfig() Config {
	var cfg Config
	flag.Parse()
	if err := cleanenv.ReadConfig(configPath(), &cfg); err != nil {
		log.Fatalf("[ERROR] read config %s", err.Error())
	}

	return cfg
}

func configPath() string {

	if *path == "" {
		*path = os.Getenv(config_path_env)
	}
	if _, err := os.Stat(*path); err == os.ErrNotExist {
		log.Fatalf("[ERROR] no such file %s", *path)
	}

	return *path
}
