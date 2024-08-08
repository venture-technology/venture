package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name      string `yaml:"name"`
	Database  Database
	Server    Server `yaml:"server-responsible"`
	Cloud     Cloud
	Cache     Cache
	Messaging Messaging
	StripeEnv StripeEnv `yaml:"stripe-env"`
}

type Server struct {
	Host   string `yaml:"host-vtx-responsible"`
	Port   int    `yaml:"port-vtx-responsible"`
	Secret string `yaml:"string-vtx-responsible"`
}

type Database struct {
	User     string `yaml:"dbuser"`
	Port     string `yaml:"dbport"`
	Host     string `yaml:"dbhost"`
	Password string `yaml:"dbpassword"`
	Name     string `yaml:"dbname"`
	Schema   string `yaml:"schema"`
}

type Cloud struct {
	Region     string `yaml:"region"`
	AccessKey  string `yaml:"accesskey"`
	SecretKey  string `yaml:"secretkey"`
	Token      string `yaml:"token"`
	Source     string `yaml:"source"`
	BucketName string `yaml:"bucketName"`
}

type Cache struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Messaging struct {
	Brokers   string `yaml:"broker"`
	Topic     string `yaml:"topic"`
	Partition int    `yaml:"partition"`
}

type StripeEnv struct {
	PublicKey string `yaml:"stripe_publickey"`
	SecretKey string `yaml:"stripe_secretkey"`
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
	return config
}
