// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"

	"github.com/matthieukhl/rgvr/internal/config"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save your Ringover API key locally",
	Long: `Prompts for your Ringover API key and stores it in
~/.config/rgvr/config.yaml so you don't need to pass it on every command.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		configPath, err := cfg.Write()
		if err != nil {
			return err
		}

		fmt.Printf("API key saved to %s\n", configPath)
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
}
