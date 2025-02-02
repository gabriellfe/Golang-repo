package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func SetupServer() {
	env := os.Getenv("ENV")
	os.Setenv("TZ", "America/Sao_Paulo")

	if env == "" {
		env = "dev"
	}

	godotenv.Load(".env." + env)
	setupConfigs(env)
	appName := os.Getenv("APP_NAME")
	LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))
}

func Refresh() {
	appName := os.Getenv("APP_NAME")
	LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))
}

func setupConfigs(env string) {
	configServer := os.Getenv("CONFIG_SERVER")
	// Read command line flags
	profile := flag.String("profile", env, "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", configServer, "Address to config server")
	configBranch := flag.String("configBranch", "main", "git branch to fetch configuration from")
	flag.Parse()

	// Pass the flag values into viper.
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func LoadConfigurationFromBranch(configServerUrl string, appName string, profile string, branch string) {
	url := fmt.Sprintf("%s/%s/%s/%s", configServerUrl, appName, profile, branch)
	slog.Info(fmt.Sprintf("Loading config from %s\n", url))
	body, err := fetchConfiguration(url)
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	parseConfiguration(body)
}

// Make HTTP request to fetch configuration from config server
func fetchConfiguration(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// Pass JSON bytes into struct and then into Viper
func parseConfiguration(body []byte) {
	var cloudConfig springCloudConfig
	err := json.Unmarshal(body, &cloudConfig)
	if err != nil {
		panic("Cannot parse configuration, message: " + err.Error())
	}
	if len(cloudConfig.PropertySources) == 0 {
		slog.Info(fmt.Sprintf("No config to load configuration for service %s\n", viper.GetString("APP_NAME")))
	}
	for key, value := range cloudConfig.PropertySources[0].Source {
		viper.Set(key, value)
		//fmt.Printf("Loading config property %v => %v\n", key, value)
	}
	if viper.IsSet("APP_NAME") {
		slog.Info(fmt.Sprintf("Successfully loaded configuration for service %s\n", viper.GetString("APP_NAME")))
	}
}

// Structs having same structure as response from Spring Cloud Config
type springCloudConfig struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	PropertySources []propertySource `json:"propertySources"`
}

type propertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}
