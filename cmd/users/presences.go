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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// presencesCmd represents the presences command
var presencesCmd = &cobra.Command{
	Use:   "presences <user_id>",
	Short: "Retrieves the current presence status for a specific user",
	Long: `Retrieves the current presence status for a specific user — whether they are available,
on a call, in snooze mode, or offline. The presence data includes the user's current availability state,
active snooze information (if any), and the device they are connected from.

Permission:
	
	Users Read required.

Monitoring impact:

	OFF: Returns presence only if userId matches your own user ID. Requesting another user's presence returns 404.
	ON: Returns any user's presence in the team. For a bulk view of all team presences, use GET /presences instead (also requires Monitoring ON).`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]
		path := fmt.Sprintf("/users/%s/presences", userID)

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		start := time.Now()
		resp, err := client.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		duration := time.Since(start)

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var presence models.Presence

		if err = json.NewDecoder(resp.Body).Decode(&presence); err != nil {
			return fmt.Errorf("decoding presence information: %w", err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.Presence{presence}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, presence); err != nil {
				return err
			}
		}

		if err := flags.IsVerbose(cmd, resp, duration); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	UsersCmd.AddCommand(presencesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presencesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presencesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
