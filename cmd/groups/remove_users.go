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

package groups

import (
	"fmt"
	"strings"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/convert"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// removeUserCmd represents the removeUser command
var removeUserCmd = &cobra.Command{
	Use:   "remove-users <group_id> <user_id1> [user_id2 ...]",
	Short: "Removes one or more users from a call group.",
	Long: `Removes one or more users from a call group.
Once removed, these users will no longer receive calls routed to this group.
If a user ID is not a member of the group, it is silently ignored.

Permission:

	Users Write required.

Monitoring:

	Required. Group membership management is a supervisory operation.
	Returns 401 Unauthorized if Monitoring is OFF on your API key.`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := args[0]
		userIDs := args[1:]

		parsedUserIDs, err := convert.StringIDsToJSON(userIDs)
		if err != nil {
			return fmt.Errorf("parsing user IDs: %w", err)
		}

		path := fmt.Sprintf("/groups/%s/users", groupID)

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Delete(path, parsedUserIDs)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		fmt.Printf("Successfully removed user(s) %s from group %s\n", strings.Join(userIDs, ", "), groupID)

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	groupsCmd.AddCommand(removeUserCmd)
}
