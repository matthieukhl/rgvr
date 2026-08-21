// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package users

import (
	"encoding/json"
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
	Use:   "get <user_id>",
	Args:  cobra.ExactArgs(1),
	Short: "Retrieves detailed information about a specific user.",
	Long: `
(Description from Ringover's official API documentation)

Retrieves detailed information about a specific user, including their profile,
assigned phone numbers, status (active, archived), and configuration.
Use this to inspect a user's setup before making changes or to verify user existence.

Permission:
	Users Read required.

Monitoring impact:
	OFF: Returns data only if userId matches your own user ID. Requesting another user's ID returns 404.
	ON: Returns data for any user in the team.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]

		path := fmt.Sprintf("/users/%s", userID)
		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err = flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var user models.User
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return fmt.Errorf("decoding user information: %w", err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.User{user}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, user); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	usersCmd.AddCommand(getCmd)
}
