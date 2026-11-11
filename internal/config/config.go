package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	DbURL				string		`json:"db_url"`
	CurrentUsername		string		`json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"
//Read
//	Takes no arguments
//	Returns the Config struct and error
//	Sets the result of getConfigFilePath and calls os.ReadFile() on that filepath.
//	The result of which is provided to json.Unmarshal()
func Read() (Config,error) {
	
	
	configFilepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error retrieving config file path: %w", err)
	}
	content, err := os.ReadFile(configFilepath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %s", err)
	}
	var cfg Config 
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse JSON data: %w", err )
	}
	return cfg, nil
	
}

// NOTE: getConfigFilePath() is a helper function to get the home dir and concatenate
// it with the global config variable configFileName
func getConfigFilePath() (string, error) {
	var result string
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving home directory %w", err)
	}
	result = home + configFileName

	return result, nil
}


// NOTE: SetUser sets current_user_name then calls write()
func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("error writing to json file: %w",err)
	}
	return nil
}

// NOTE: write marshals cfg to JSON and writes to the file
func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshalling json data %w: ", err)
	}  
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error retrieving config file path: %w", err)
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}
