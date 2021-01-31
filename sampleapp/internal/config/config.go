package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/poncos/go-things/sampleapp/internal/utils"

	"gopkg.in/yaml.v2"
)

const configFile string = "conf/config.yml"

// The SampleAppConfig is the main data type representing the app global config
type SampleAppConfig struct {
	LogLevel  string `yaml:"logLevel"`

	OutputInfo struct {
		Workspace string `yaml:"workspace"`
		OutputFile string `yaml:"outputFile"`
	} `yaml:"outputInfo"`

	AppInfo struct {
		Version  string `yaml:"version"`
		Description string `yaml:"description"`
	} `yaml:"appInfo"`

	ConfigDir string
}

// LoadConfig read the global configuration from the default location
func LoadConfig() SampleAppConfig {

	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	exPath := filepath.Dir(ex)

	slashPath := fmt.Sprintf("%s/%s", exPath, configFile)
	sampleAppConfig := LoadConfigFromFile(slashPath)

	sampleAppConfig.ConfigDir = fmt.Sprintf("%s/conf", exPath)

	return sampleAppConfig
}

// LoadConfigFromDir reads the global configuration and returns the data type with the information
func LoadConfigFromFile(configFile string) SampleAppConfig {

	utils.Debug("LoadConfigFromFile %s", configFile)

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return unmarshallGlobalConfig(b)
}

func unmarshallGlobalConfig(data []byte) SampleAppConfig {
	appConfig := SampleAppConfig{}

	err := yaml.Unmarshal([]byte(data), &appConfig)
	if err != nil {
		log.Fatal(err)
	}

	return appConfig
}
