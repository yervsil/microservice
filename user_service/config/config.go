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
		Server      Server
		Mongo 		MongoConfig
		Redis 		RedisConfig
	}


	MongoConfig struct {
		URI      string `mapstructure:"MONGO_URI"`
		User     string `mapstructure:"MONGO_USER"`
		Password string
		Name     string `mapstructure:"databaseName"`
	}

	RedisConfig struct {
		RedisAddr      string `mapstructure:"RedisAddr"`
		RedisPassword  string `mapstructure:"RedisPassword"`
		RedisDB        string `mapstructure:"RedisDB"`
		RedisDefaultdb string `mapstructure:"RedisDefaultdb"`
		MinIdleConns   int    `mapstructure:"MinIdleConns"`
		PoolSize       int    `mapstructure:"PoolSize"`
		PoolTimeout    int    `mapstructure:"PoolTimeout"`
		Password       string `mapstructure:"Password"`
		DB             int    `mapstructure:"DB"`
	}

	Server struct {
		Port              string		`mapstructure:"Port"`
		Timeout           time.Duration `mapstructure:"Timeout"`
		ReadTimeout       time.Duration `mapstructure:"ReadTimeout"`
		WriteTimeout      time.Duration `mapstructure:"WriteTimeout"`
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

	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return nil, err 
	}

	if err := viper.UnmarshalKey("mongodb", &cfg.Mongo); err != nil {
		return nil, err 
	}

	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
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

	if err := viper.BindEnv("MONGO_PASS"); err != nil {
		return err 
	}

	return nil 
}