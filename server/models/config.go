package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ENV_KEY string = "APP_ENVIORNMENT"
const ENV_DEV string = "development"
const ENV_TEST string = "test"
const ENV_PROD string = "production"

const DEFAULT_CONFIG_PATH string = "./server.config.json"

// GetEnv retrieves app environment from environment variable key in ENV_KEY, defaults to development. Returns empty
// string for error
func GetEnv() (error, string) {
	envVal := os.Getenv(ENV_KEY)
	if envVal == "" {
		envVal = ENV_DEV
	}

	if envVal != ENV_DEV && envVal != ENV_TEST && envVal != ENV_PROD {
		return fmt.Errorf("Unknown environment: \"%s\"", envVal), ""
	}

	return nil, envVal
}

type Config struct {
	Port int

	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string

	FCMServerKey string
}

// LoadFile loads a specified json file and retrieves Config value for app environmnet
func LoadConfigFile(path string) (error, *Config) {
	// Read file
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e, nil
	}

	// Unmarshal
	type ConfigFileType map[string]*Config
	var data ConfigFileType

	json.Unmarshal(file, &data)

	// Check app env data exists in file
	err, env := GetEnv()
	if err != nil {
		return err, nil
	}

	if _, ok := data[env]; !ok {
		// If no key for current app env exists in loaded file, error
		return fmt.Errorf("No config for app environment \"%s\" found in file", env), nil
	}

	return nil, data[env]
}

// GetDbConnOpts generates PostGreSQL connection options for the current configuration
func (c *Config) GetDbConnOpts(extraOpts string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s %s",
		c.DbHost,
		c.DbUser,
		c.DbPassword,
		c.DbName,
		extraOpts,
	)
}
