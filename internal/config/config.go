package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:4001"`
	Timeout     time.Duration `yaml:"timeout"  env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout"  env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true"`
}

func MustConfig() *Config {
	//configPath := os.Getenv("CONFIG_PATH")
	//if configPath == "" {
	//	log.Fatal("CONFIG_PATH is not set")
	//}
	////log.Println(configPath)
	//
	//if _, err := os.Stat(configPath); os.IsNotExist(err) {
	//	log.Fatalf("config file does not exist: %s", configPath)
	//}

	var cfg Config
	a := "./config/local.yaml"

	if err := cleanenv.ReadConfig(a, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
