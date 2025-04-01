package config

import "github.com/spf13/viper"

type Config struct {
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`
    DR600 struct {
        BaseURL string `mapstructure:"base_url"`
    } `mapstructure:"dr600"`
}

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigFile(path)
    viper.SetConfigType("yaml")
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}