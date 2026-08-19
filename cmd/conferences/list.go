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

package conferences

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves all conference rooms (audio bridges) configured for your team.",
	Long: `
Retrieves all conference rooms (audio bridges) configured for your team.
A conference room allows multiple participants to join a shared audio call
using a dedicated phone number and PIN code. Each conference includes its assigned number,
participant PIN, admin PIN, and capacity settings. The total count is in the list_count field.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Required. Conference management is a team-level supervisory feature. Returns 401 Unauthorized if Monitoring is OFF on your API key.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		conferences, reqInfo, err := httpClient.GetConferences()
		if err != nil {
			return err
		}

		if conferences == nil {
			fmt.Println("No conference rooms found.")
			return nil
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		switch format {
		case "table":
			if err := formats.Table(os.Stdout, conferences.List); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}

		default:
			if err := formats.JSON(os.Stdout, conferences.List); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil

	},
}

func init() {
	conferencesCmd.AddCommand(listCmd)
	listCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
