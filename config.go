package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const CONFIG_PATH string = "config.json"

func saveConfig(key string, value string) {
	config[key] = value

	// encode to json
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	// save to file
	file, err := os.OpenFile(CONFIG_PATH, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(jsonConfig)

	fmt.Println(key, " data saved at: ", CONFIG_PATH)
}

func getConfig() map[string]string {
	// Create config file if it doesn't exist
	file, err := os.OpenFile(CONFIG_PATH, os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonConfig, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]string
	if len(jsonConfig) == 0 {
		return data
	}

	err = json.Unmarshal(jsonConfig, &data)
	if err != nil {
		log.Fatal(err)
	}

	config = data

	return data
}

func getConfigValue(key string) string {
	// search key in object
	value, ok := config[key]
	if !ok {
		log.Fatalf("Key %s not found in saved data: %s", key, CONFIG_PATH)
	}
	return value
}
