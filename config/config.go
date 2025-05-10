package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/pkg/stringcommon"
	"github.com/venture-technology/venture/pkg/utils"
)

const (
	ServerEnvironment = "ENVIRONMENT"
)

func LoadServerEnvironmentVars(service, serverEnv string) error {
	if stringcommon.Empty(serverEnv) || serverEnv == "development" {
		viper.SetDefault(ServerEnvironment, "development")
	}

	if viper.GetString(ServerEnvironment) == "development" {
		viper.SetConfigType("json")
		viper.SetConfigName(viper.GetString(ServerEnvironment)) // development

		path, err := utils.GetAbsPath()
		if err != nil {
			return err
		}

		viper.AddConfigPath(path)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Failed to read config file:", err)
		}

		viper.AutomaticEnv()
	} else {
		viper.AutomaticEnv()
	}

	return nil
}

func DevEnv() bool {
	return viper.GetString(ServerEnvironment) == "development"
}

func ProdEnv() bool {
	return viper.GetString(ServerEnvironment) == "production"
}

func LoadStaticFile(pathToRoot, filename string) ([]byte, error) {
	_, path, _, _ := runtime.Caller(1)
	fullPath := filepath.Join(path, pathToRoot, "data", filename)
	return ioutil.ReadFile(fullPath)
}
