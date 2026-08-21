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

package users

import (
	"fmt"
	"strconv"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive <user_id>",
	Short: "Permanently archives a user by their unique identifier.",
	Long: `
Permanently archives a user by their unique identifier.
Archiving a user disables their account but preserves their data for future reactivation.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Not needed.

Warning — Destructive operation:

	This route does not enforce any specific permission category.
	Any valid API key can archive any user in the team.
	Exercise extreme caution when using this endpoint, especially in automated workflows.
	Consider restricting API key distribution accordingly.
`,
	Args: cobra.ExactArgs(1),
	Example: `
  # Archive a user with ID 12345
  rgvr users archive 12345
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			return fmt.Errorf("invalid user ID: %s. User ID must be a valid integer", userID)
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		reqInfo, err := httpClient.DeleteUser(userIDInt, cmd.Name())
		if err != nil {
			return fmt.Errorf("failed to archive user with ID %d: %v", userIDInt, err)
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		fmt.Printf("User with ID %d has been archived successfully.\n", userIDInt)
		return nil
	},
}

func init() {
	usersCmd.AddCommand(archiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// archiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// archiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
