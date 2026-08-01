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
	"time"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/convert"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// addUserCmd represents the addUser command
var addUserCmd = &cobra.Command{
	Use:   "add-users <group_id> <user_id1> [<user_id2> ...]",
	Short: "Adds a user to a call group.",
	Long: `Adds a user to a call group. Once added, the user will receive calls routed
to this group according to the group's ring strategy (simultaneous, round-robin, etc.).
The user must exist in the team. If the user is already a member of the group, the request has no effect.

Permission:

	Users Write required.

Monitoring:

	Required. Group membership management is a supervisory operation. Returns 401 Unauthorized if Monitoringis OFF on your API key.`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := args[0]
		userIDs := args[1:]

		parsedUserIDs, err := convert.StringIDsToJSON(userIDs)
		if err != nil {
			return fmt.Errorf("parsing user IDs: %w", err)
		}

		path := fmt.Sprintf("/groups/%s/users", groupID)

		client := cmd.Context().Value(internal.ClientContextKey).(*internal.Client)

		start := time.Now()
		resp, err := client.Post(path, parsedUserIDs)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		duration := time.Since(start)

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		printSuccess(groupID, userIDs)

		if err := flags.IsVerbose(cmd, resp, duration); err != nil {
			return err
		}

		return nil
	},
}

func printSuccess(groupID string, userIDs []string) {
	if len(userIDs) == 1 {
		fmt.Printf("Successfully added user %s to group %s\n", userIDs[0], groupID)
	} else {
		userIdsStr := strings.Join(userIDs, ", ")
		fmt.Printf("Successfully added users %s to group %s\n", userIdsStr, groupID)
	}
}

func init() {
	groupsCmd.AddCommand(addUserCmd)
}
