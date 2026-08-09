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

package numbers

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <number>",
	Short: "Retrieves detailed information about a specific phone number.",
	Long: `Retrieves detailed information about a specific phone number,
including its assignment (user, IVR, conference), country, capabilities,
and configuration. The number is identified by its phone number in E.164 format
without the + prefix (e.g., 33140000000).

Permission:

	Numbers Read required.

Monitoring impact:

	OFF: Returns the number only if it is assigned to you.
	Requesting a number belonging to another user returns 404.
	ON: Returns details for any number in the team.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		number := args[0]

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		numberResponse, reqInfo, err := httpClient.GetNumber(number)
		if err != nil {
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		switch format {
		case "table":
			if err := formats.Table(os.Stdout, []models.Number{*numberResponse}); err != nil {
				return err
			}
		default:
			if err := formats.JSON(os.Stdout, []models.Number{*numberResponse}); err != nil {
				return err
			}
		}

		return nil

	},
}

func init() {
	numbersCmd.AddCommand(getCmd)
	getCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
