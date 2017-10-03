package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// EnvKey holds the name of the environment variable which will hold the
// application environment.
//
// The terms "environment" & "environment variable" are very similar terms used
// to designate different things:
//
//     - environment variable = Variable provided by shell
//     - environment = where application is running (ex., production, staging, test)
const EnvKey string = "APP_ENVIORNMENT"

// EnvDev holds the environment value to signify the application is running
// for the purpose of development
const EnvDev string = "development"

// EnvTest holds the environment value to signify the application is running for
// the purpose of teting
const EnvTest string = "test"

// EnvProd holds the environment value to signify the application is running to
// server production traffic
const EnvProd string = "production"

// DefaultConfigPath is the path where the server config.json file is located
// by default.
const DefaultConfigPath string = "./server.config.json"

// GetEnv retrieves app environment from environment variable key in ENV_KEY,
// defaults to development. Returns empty string for error.
func GetEnv() (string, error) {
	envVal := os.Getenv(EnvKey)
	if envVal == "" {
		envVal = EnvDev
	}

	if envVal != EnvDev && envVal != EnvTest && envVal != EnvProd {
		return fmt.Errorf("Unknown environment: \"%s\"", envVal), ""
	}

	return nil, envVal
}

// Config holds all application configuration information
type Config struct {
	// Port is the network port to listen on for requests
	Port int

	// DbHost is the host to access the database at
	DbHost string

	// DbUser is the username to access the database with
	DbUser string

	// DbPassword is the password to access the database with
	DbPassword string

	// DbName is the name of the database to store information in
	DbName string

	// FCMServerKey is the Firebase Cloud Messaging key to provide for
	// API authentication
	FCMServerKey string
}

// LoadConfigFile loads a specified json file and retrieves Config value for
// app environmnet
func LoadConfigFile(path string) (*Config, error) {
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
	env, err := GetEnv()
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
