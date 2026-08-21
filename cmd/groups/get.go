// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package groups

import (
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
	Use:   "get <group_id>",
	Short: "Retrieves detailed information about a specific call group.",
	Long: `
Description:

	Retrieves detailed information about a specific call group, including its members,
	ring strategy, ring durations, and IVR assignments. Use this to inspect a group's configuration
	before making changes (e.g., adding/removing members, modifying ring duration).

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring impact:

	OFF: Returns the group only if you are a member of it. Returns 404 otherwise.
	ON: Returns any group in the team regardless of your membership.

Member pagination: 

	'--limit' and '--offset' paginate only the users sub-array (group members) of the response —
	they do not affect the group object itself. The response does not include a total member count:
	keep incrementing limit_offset until an empty users array is returned.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupID := args[0]

		params, err := flags.BuildPaginationParams(cmd)
		if err != nil {
			return err
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		group, reqInfo, err := httpClient.GetGroupByID(groupID, params)
		if err != nil {
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.Group{*group}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, group); err != nil {
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
	groupsCmd.AddCommand(getCmd)
	getCmd.Flags().IntP("limit", "l", 0, "Maximum number of results to return per page. Default: server-defined.")
	getCmd.Flags().Int("offset", 0, "Number of results to skip for pagination. Default: 0.")
}
