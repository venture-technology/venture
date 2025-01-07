package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name              string            `yaml:"name"`
	Database          Database          `yaml:"database"`
	Server            Server            `yaml:"server"`
	Cloud             Cloud             `yaml:"cloud"`
	Cache             Cache             `yaml:"cache"`
	Uchiha            Uchiha            `yaml:"uchiha"`
	Mongo             Mongo             `yaml:"mongo"`
	GoogleCloudSecret GoogleCloudSecret `yaml:"google-cloud-secret"`
	StripeEnv         StripeEnv         `yaml:"stripe-env"`
	Admin             Admin             `yaml:"admin"`
}

type Server struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Secret string `yaml:"string"`
}

type Database struct {
	User     string `yaml:"dbuser"`
	Port     string `yaml:"dbport"`
	Host     string `yaml:"dbhost"`
	Password string `yaml:"dbpassword"`
	Name     string `yaml:"dbname"`
}

type Cloud struct {
	Region     string `yaml:"region"`
	AccessKey  string `yaml:"accesskey"`
	SecretKey  string `yaml:"secretkey"`
	Token      string `yaml:"token"`
	Source     string `yaml:"source"`
	BucketName string `yaml:"bucketname"`
}

type Cache struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Uchiha struct {
	Address string `yaml:"address"`
	Queue   string `yaml:"queue"`
}

type Mongo struct {
	Address    string `yaml:"address"`
	Database   string `yaml:"dbname"`
	Collection string `yaml:"collection"`
}

type GoogleCloudSecret struct {
	ApiKey                 string `yaml:"apikey"`
	EndpointMatrixDistance string `yaml:"endpoint-matrix-distance"`
}

type StripeEnv struct {
	PublicKey string `yaml:"publickey"`
	SecretKey string `yaml:"secretkey"`
}

type Admin struct {
	ApiKey string `yaml:"apikey"`
	Port   string `yaml:"port"`
}

var config *Config

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	config = &conf
	return config, nil
}

func Get() *Config {
	if config == nil {
		config, err := Load("../../../config/config.yaml")
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		return config
	}
	return config
}
