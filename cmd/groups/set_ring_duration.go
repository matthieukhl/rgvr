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
	"strconv"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// setRingDurationCmd represents the setRingDuration command
var setRingDurationCmd = &cobra.Command{
	Use:   "set-ring-duration <group_id> <user_id> <ring_duration>",
	Short: "Sets how long (in seconds) the phone rings for a specific user within a call group before the call is forwarded to the next member or dropped.",
	Long: `
Description:

	Sets how long (in seconds) the phone rings for a specific user within a call group before the call is forwarded to the next member or dropped.
	This is a per-user, per-group setting — a user can have different ring durations in different groups.
	'ring_duration' must be between 0 and 295 and a multiple of 5".

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Required. Group configuration changes are supervisory operations. Returns 401 Unauthorized if Monitoring is OFF on your API key.`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := args[0]
		userID := args[1]
		ringDuration, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("invalid ring duration, must be an integer must be between 0 and 295 and a multiple of 5: %w", err)
		}

		if (ringDuration < 0 || ringDuration > 295) || ringDuration%5 != 0 {
			return fmt.Errorf("invalid ring duration, must be an integer between 0 and 295 and a multiple of 5")
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		reqInfo, err := httpClient.PatchGroupRingduration(groupID, userID, ringDuration)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		fmt.Printf("Successfully set ring duration for user %s in group %s to %d seconds.\n", userID, groupID, ringDuration)

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	groupsCmd.AddCommand(setRingDurationCmd)
}
