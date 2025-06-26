package main

import (
	"os"
	"strconv"
)

func GetString(Key, fallback string) string {
	val, ok := os.LookupEnv(Key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(Key string, fallback int) int {
	val, ok := os.LookupEnv(Key)
	if !ok {
		return fallback
	}
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
