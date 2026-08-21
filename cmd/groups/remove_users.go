// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
	Long: `
Description:

	Removes one or more users from a call group.
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
