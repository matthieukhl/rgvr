// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of your Ringover API key",
	Long: `Checks the status of your Ringover API key by making
a test request to the API. If the API key is valid, it will return a success message.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient, err := client.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		// Make a test request to the API to check if the API key is valid
		path := "/users?limit_count=1" // A simple endpoint that requires authentication as Ringover's public API does not have a dedicated endpoint for checking API key validity

		resp, reqInfo, err := httpClient.Get(path)
		if err != nil {
			return fmt.Errorf("making test request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("API key is invalid. Please run `rgvr auth login` to set a valid API key")
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		fmt.Println("API key is valid.")
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(statusCmd)
}
