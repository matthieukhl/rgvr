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

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// toggleAccessCmd represents the toggleAccess command
var toggleAccessCmd = &cobra.Command{
	Use:   "toggle-access <group_id>",
	Short: "Toggles the 'is_jumper' flag on a specific group.",
	Long: `Toggles the is_jumper flag on a specific group.
When enabled (is_jumper = true), users can freely join
and leave this group without requiring an administrator action.
When disabled, group membership is managed exclusively by administrators.
This corresponds to the "Free access" switch in the Ringover dashboard.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Required. This is a team-level configuration change.
	Returns 401 Unauthorized if Monitoring is OFF on your API key.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := args[0]

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		reqInfo, err := client.PatchGroupAccess(groupID)
		if err != nil {
			return err
		}

		group, _, err := client.GetGroupByID(groupID, nil)

		fmt.Printf("Group %q (ID: %d): free access is now %t\n", group.Name, group.GroupID, group.IsJumper)

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	groupsCmd.AddCommand(toggleAccessCmd)
}
