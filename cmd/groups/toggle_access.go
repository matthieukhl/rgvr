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

		path := fmt.Sprintf("/groups/%s/isjumper/toggle", groupID)

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Patch(path, nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	groupsCmd.AddCommand(toggleAccessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleAccessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleAccessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
