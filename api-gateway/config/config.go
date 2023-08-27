package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


type (
	Config struct {
		Env 		string 		`mapstructure:"env"`
		HTTP        HTTPConfig
	}


	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

)



func InitConfig() (*Config, error) {
	//viper.AddConfigPath("../../config")
	// viper.SetConfigName("main")
    viper.AddConfigPath("./config")
	viper.SetConfigName("main")

	
	if err := viper.ReadInConfig(); err != nil {
		return nil, err 
	}

	var cfg Config
	if err := viper.UnmarshalKey("env", &cfg.Env); err != nil {
		return nil, err 
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return nil, err 
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err 
	}

	return &cfg, nil 
}

func parseEnv(cfg *Config) error {
	//err := godotenv.Load("../../.env")
	err := godotenv.Load("./.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	// if err := viper.BindEnv("MONGO_PASS"); err != nil {
	// 	return err 
	// }

	return nil 
}