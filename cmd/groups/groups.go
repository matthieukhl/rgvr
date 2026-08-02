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
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Retrieves call groups (ring groups) configured for your team.",
	Long: `
Description:
	
	Retrieves call groups (ring groups) configured for your team.
	Groups define how incoming calls are distributed among a set of users (simultaneous ring, round-robin, etc.).
	Each group includes its member list, ring strategy, and configuration.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring impact:

	OFF: Returns only the groups you belong to as a member.
	ON: Returns all groups in the team, regardless of your membership.

Pagination: 

	Results are returned in a stable order (by group id), so the same offset always yields the same page.
	To read every group, keep limit_count fixed and increase limit_offset by that amount on each request — 
	this avoids the duplicates and gaps you would get from an unordered list.`,
	Args: cobra.NoArgs,
	PersistentPreRunE: func(cd *cobra.Command, args []string) error {
		httpClient, err := client.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		ctx := context.WithValue(cd.Context(), client.ClientContextKey, httpClient)
		cd.SetContext(ctx)
		return nil
	},
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/groups"

		params, err := flags.BuildPaginationParams(cd)
		if err != nil {
			return err
		}

		if len(params) > 0 {
			path += "?" + params.Encode()
		}

		client := cd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var groupsResponse models.ListResponse[models.Group]

		if err = json.NewDecoder(resp.Body).Decode(&groupsResponse); err != nil {
			return fmt.Errorf("decoding groups information: %w", err)
		}

		format, err := cd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, groupsResponse.List); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, groupsResponse); err != nil {
				return err
			}
		}

		if err := flags.IsVerbose(cd, reqInfo); err != nil {
			return err
		}

		return nil

	},
}

func init() {
	cmd.RootCmd.AddCommand(groupsCmd)
	groupsCmd.Flags().String("format", "json", "Choose the output's format: table / json")

	groupsCmd.Flags().IntP("limit", "l", 0, "Maximum number of results to return per page. Default: server-defined.")
	groupsCmd.Flags().Int("offset", 0, "Number of results to skip for pagination. Default: 0.")
}
