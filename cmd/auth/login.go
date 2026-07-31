/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah <matthieu.khairallah@proton.me>

	    This program is free software: you can redistribute it and/or modify
	    it under the terms of the GNU Affero General Public License as published by
	    the Free Software Foundation, either version 3 of the License, or
	    (at your option) any later version.

	    This program is distributed in the hope that it will be useful,
	    but WITHOUT ANY WARRANTY; without even the implied warranty of
	    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	    GNU Affero General Public License for more details.

	    You should have received a copy of the GNU Affero General Public License
	    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save your Ringover API key locally",
	Long: `Prompts for your Ringover API key and stores it in
~/.config/rgvr/config.yaml so you don't need to pass it on every command.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		apiKey, err := registerAPIKey()
		if err != nil {
			return err
		}

		region, err := registerRegion()
		if err != nil {
			return err
		}

		configDir, err := internal.GetConfigDir()
		configPath := filepath.Join(configDir, "config.yaml")

		if err := os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("creating config directory: %w", err)
		}

		viper.Set("api_key", apiKey)
		viper.Set("region", region)

		if err := viper.WriteConfigAs(configPath); err != nil {
			return fmt.Errorf("writing config file: %w", err)
		}

		if err := os.Chmod(configPath, 0600); err != nil {
			return fmt.Errorf("setting config file permissions: %w", err)
		}

		fmt.Printf("API key saved to %s\n", configPath)
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
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

func registerRegion() (string, error) {
	fmt.Print("Please enter your Ringover region ('eu' or 'us'): ")
	var region string
	_, err := fmt.Scanln(&region)
	if err != nil {
		return "", fmt.Errorf("reading region: %w", err)
	}

	if region != "eu" && region != "us" {
		return "", fmt.Errorf("invalid region: %s. Please enter 'eu' or 'us'", region)
	}
	return region, nil
}
