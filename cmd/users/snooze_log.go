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

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

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

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
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
