package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"strconv"
)

func ReadConfigString(filePath string, section string, key string) string {
	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return cfg.Section(section).Key(key).String()
}

func ReadConfigUInt64(filePath string, section string, key string) (uint64, error) {
	s := ReadConfigString(filePath, section, key)

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return i, err
	}

	return i, nil
}

func ReadConfigInt(filePath string, section string, key string) (int, error) {
	n, err := ReadConfigUInt64(filePath, section, key)
	return int(n), err
}

func ReadConfigFloat64(filePath string, section string, key string) (float64, error) {
	s := ReadConfigString(filePath, section, key)

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return i, err
	}

	return i, nil
}
