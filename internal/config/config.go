// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/term"
)

type Region int

const (
	EU Region = iota
	US
)

var regionName = map[Region]string{
	EU: "eu",
	US: "us",
}

func (r Region) String() string {
	return regionName[r]
}

type Config struct {
	ApiKey string
	UserID int
	Region Region
}

func (c *Config) Write() (configPath string, err error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", fmt.Errorf("fetching config directory: %w", err)
	}

	configPath = filepath.Join(configDir, "config.yaml")

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", fmt.Errorf("creating config directory: %w", err)
	}

	viper.Set("api_key", c.ApiKey)
	viper.Set("user_id", c.UserID)
	viper.Set("region", c.Region.String())

	if err := viper.WriteConfigAs(configPath); err != nil {
		return "", fmt.Errorf("writing config file: %w", err)
	}

	if err := os.Chmod(configPath, 0600); err != nil {
		return "", fmt.Errorf("setting config file permissions: %w", err)
	}

	return configPath, nil
}

func NewConfig() (*Config, error) {
	apiKey, err := registerAPIKey()
	if err != nil {
		return nil, err
	}

	userID, err := registerUserID()
	if err != nil {
		return nil, err
	}

	region, err := registerRegion()
	if err != nil {
		return nil, err
	}

	return &Config{
		ApiKey: apiKey,
		UserID: userID,
		Region: region,
	}, nil
}

func registerUserID() (int, error) {
	fmt.Print("Please enter your Ringover user ID:")

	var userID int
	_, err := fmt.Scanln(&userID)
	if err != nil {
		return 0, fmt.Errorf("reading user ID: %w", err)
	}

	return userID, nil
}

func registerAPIKey() (string, error) {
	fmt.Print("Please enter your Ringover API key: ")

	byteKey, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("reading API key: %w", err)
	}

	apiKey := strings.TrimSpace(string(byteKey))
	if apiKey == "" {
		return "", fmt.Errorf("API key cannot be empty")
	}

	return apiKey, nil
}

func registerRegion() (Region, error) {
	fmt.Print("Please enter your Ringover region ('eu' or 'us'): ")
	var region string
	_, err := fmt.Scanln(&region)
	if err != nil {
		return 0, fmt.Errorf("reading region: %w", err)
	}

	return parseRegion(region)
}

func parseRegion(s string) (Region, error) {
	for region, name := range regionName {
		if name == s {
			return region, nil
		}
	}

	return 0, fmt.Errorf("invalid region: %s. Please enter 'eu' or 'us'", s)
}
