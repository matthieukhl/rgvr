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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <group_id>",
	Short: "Retrieves detailed information about a specific call group.",
	Long: `Retrieves detailed information about a specific call group, including its members,
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

		path := fmt.Sprintf("/groups/%s", groupID)

		params, err := flags.BuildPaginationParams(cmd)
		if err != nil {
			return err
		}

		if len(params) > 0 {
			path += "?" + params.Encode()
		}

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		start := time.Now()
		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}
		defer resp.Body.Close()
		duration := time.Since(start)

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var group models.Group

		if err = json.NewDecoder(resp.Body).Decode(&group); err != nil {
			return fmt.Errorf("decoding groups information: %w", err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.Group{group}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, group); err != nil {
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
	groupsCmd.AddCommand(getCmd)
	getCmd.Flags().IntP("limit", "l", 0, "Maximum number of results to return per page. Default: server-defined.")
	getCmd.Flags().Int("offset", 0, "Number of results to skip for pagination. Default: 0.")
}
