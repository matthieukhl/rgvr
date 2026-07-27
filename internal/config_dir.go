package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetConfigDir returns the path to the configuration directory for rgvr.
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving home directory: %w", err)
	}
	configDir := filepath.Join(home, ".config", "rgvr")
	return configDir, nil
}
