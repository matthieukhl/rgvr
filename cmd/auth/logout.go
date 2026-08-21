// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out and remove your Ringover API key from local storage",
	Long: `Removes your Ringover API key from local storage by deleting
the config file (~/.config/rgvr/config.yaml).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile := viper.ConfigFileUsed()

		if configFile == "" {
			fmt.Println("Not currently logged in.")
			return nil
		}

		if err := os.Remove(configFile); err != nil {
			return fmt.Errorf("removing config file: %w", err)
		}
		fmt.Println("Successfully logged out.")
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(logoutCmd)
}
