package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	dbURL				string		`json/"db_url"`
	currentUsername		string		`json/"current_user_name"`
}

const configFileName = ".gatorconfig.json"
// TODO: Export a function that reads the JSON file found at `~/.gatorconfig.json`
// and returns a Config struct. It should read the file from the HOME direectory
// then decode the JSON string into a new Config struct. 
// NOTE: Use os.UserHomeDir to access the home directory.
func Read() (Config,error) {
	var c Config
	configFilepath := ""
	home, homeErr := os.UserHomeDir()
	if homeErr != nil {
		return Config{}, fmt.Errorf("Failed to retrieve home folder", homeErr)
	}
	configFilepath = home + configFileName
	content, err := os.ReadFile(configFilepath)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read file: %s", err)
	}
	jsonErr := json.Unmarshal(content, &c)
	if jsonErr != nil {
		return Config{}, fmt.Errorf("Failed to parse JSON data %w", jsonErr )
	}
	return c, nil
	
}

// NOTE: getConfigFilePath() is a helper function to get the home dir and concatenate
// it with the global config variable configFileName
func getConfigFilePath() (string, error) {
	var result string
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error retrieving home directory %w", err)
	}
	result = home + configFileName
	return result, nil
}


// NOTE: SetUser sets current_user_name then calls write()
func (c Config) SetUser(username string) error {

}

// NOTE: write marshals cfg to JSON and writes to the file
