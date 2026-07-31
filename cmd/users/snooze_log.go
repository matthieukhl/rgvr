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

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log <user_id>",
	Short: "Retrieves the snooze activity log for a specific user",
	Long: `Retrieves the snooze activity log for a specific user.
Each entry shows a past snooze event with the snooze type (reason),
start/end timestamps, and duration. Use this to audit a user's availability patterns over time.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Not needed. This route works regardless of the Monitoring flag.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]
		path := fmt.Sprintf("/users/%s/snooze/log", userID)

		params, err := flags.BuildPaginationParams(cmd)
		if err != nil {
			return err
		}

		if len(params) > 0 {
			path += "?" + params.Encode()
		}

		client := cmd.Context().Value(internal.ClientContextKey).(*internal.Client)

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

		var snoozeLog models.SnoozeLog

		if err = json.NewDecoder(resp.Body).Decode(&snoozeLog); err != nil {
			return fmt.Errorf("decoding snooze log information: %w", err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			// Display all snooze entries returned by the query
			if err = formats.Table(os.Stdout, snoozeLog.SnoozeLogList); err != nil {
				return err
			}

			// Display the summary of the snooze log
			if err = formats.Table(os.Stdout, []models.SnoozeLog{snoozeLog}); err != nil {
				return err
			}

		default:
			if err = formats.JSON(os.Stdout, snoozeLog); err != nil {
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
	snoozeCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	logCmd.Flags().IntP("limit", "l", 0, "Maximum number of results to return per page. Default: server-defined.")
	logCmd.Flags().Int("offset", 0, "Number of results to skip for pagination. Default: 0.")
}
