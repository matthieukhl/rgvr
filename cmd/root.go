// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "rgvr",
	Version: appVersion,
	Short:   "A CLI to interact with Ringover's public API.",
	Long: `A CLI to interact with Ringover's public API.
It allows you to manage your Ringover account directly from the command line.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	RootCmd.Version = appVersion
	RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String("api-key", "", "Your Ringover API key set by using the `rgvr auth login` command")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Display additional information about command execution")
}

func initConfig() {
	viper.SetEnvPrefix("RGVR")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configDir, err := config.GetConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting config directory: %v\n", err)
		os.Exit(1)
	}

	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if !errors.As(err, &notFoundErr) {
			fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := viper.BindPFlag("api_key", RootCmd.PersistentFlags().Lookup("api-key")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding API key flag: %v\n", err)
		os.Exit(1)
	}
}
