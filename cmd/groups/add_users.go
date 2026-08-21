// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package groups

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/convert"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// addUserCmd represents the addUser command
var addUserCmd = &cobra.Command{
	Use:   "add-users <group_id> <user_id1> [<user_id2> ...]",
	Short: "Adds a user to a call group.",
	Long: `
Description:

	Adds a user to a call group. Once added, the user will receive calls routed
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

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Post(path, parsedUserIDs)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("unexpected response from API (%s): %s", resp.Status, string(bodyBytes))
		}

		fmt.Printf("Successfully added user(s) %s to group %s\n", strings.Join(userIDs, ", "), groupID)

		return nil
	},
}

func init() {
	groupsCmd.AddCommand(addUserCmd)
}
