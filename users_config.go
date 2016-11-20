package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadUsersFile reads and parses the users file
func ReadUsersFile(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read users config file: %s", err)
	}

	var config usersFile
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse users config file: %s", err)
	}

	users := make(map[string]string)

	for _, u := range config.Users {
		users[u.ID] = u.Username
	}

	return users, nil
}

type usersFile struct {
	Users []struct {
		ID       string `yaml:"id"`
		Username string `yaml:"username"`
	} `yaml:"users"`
}
