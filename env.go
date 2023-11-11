package main

import "os"

func setEnvVariable(key string, value string) string {
	os.Setenv(key, value)
	return os.Getenv(key)
}
