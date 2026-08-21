// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/matthieukhl/rgvr/cmd"
	"github.com/spf13/cobra"
)

// AuthCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login, logout, and manage your Ringover API key",
	Long:  `The auth command allows you to log in, log out, and manage your Ringover API key.`,
}

func init() {
	cmd.RootCmd.AddCommand(AuthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	AuthCmd.PersistentFlags().String("login", "", "Register your Ringover API key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AuthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
