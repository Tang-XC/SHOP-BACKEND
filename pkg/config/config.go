package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	Env                    string `yaml:"env"`
	Address                string `yaml:"address"`
	Port                   int    `yaml:"port"`
	GracefulShutdownPeriod int    `yaml:"gracefulShutdownPeriod"`
	JwtSecret              string `yaml:"jwtSecret"`
}

type DBConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"'`
	Secure          bool   `yaml:"secure"`
	BucketName      string `yaml:"bucketName"`
}
type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Minio  MinioConfig  `yaml:"minio"`
}

func Parse(appConfig string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(appConfig)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
