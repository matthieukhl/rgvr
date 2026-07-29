/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah

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
	"net/http"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of your Ringover API key",
	Long: `Checks the status of your Ringover API key by making
a test request to the API. If the API key is valid, it will return a success message.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		// Make a test request to the API to check if the API key is valid
		path := "/users?limit_count=1" // A simple endpoint that requires authentication as Ringover's public API does not have a dedicated endpoint for checking API key validity

		resp, err := client.Get(path)
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

		fmt.Println("API key is valid.")
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(statusCmd)
}
