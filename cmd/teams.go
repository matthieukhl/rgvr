/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah

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

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Retrieve team information.",
	Long:  `Retrieves a complete team object containing lists of numbers, users, ivrs, conferences, tags and groups.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		client, err := internal.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		ctx := context.WithValue(cmd.Context(), clientContextKey, client)
		cmd.SetContext(ctx)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "/teams"
		client := cmd.Context().Value(clientContextKey).(*internal.Client)

		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving team information: %w", err)
		}
		defer resp.Body.Close()

		var team models.Team
		if err = json.NewDecoder(resp.Body).Decode(&team); err != nil {
			return fmt.Errorf("decoding team information: %w", err)
		}

		table := tablewriter.NewTable(os.Stdout, tablewriter.WithHeaderAutoFormat(tw.Off))

		table.Header([]string{
			"Team ID",
			"Name",
			"Total Numbers Count",
			"Total Users Count",
			"Total Conferences Count",
			"Total IVRs Count",
			"Total Tags Count",
			"Total Groups Count",
		})

		table.Append([]string{
			fmt.Sprintf("%d", team.TeamID),
			team.Name,
			fmt.Sprintf("%d", team.TotalNumbersCount),
			fmt.Sprintf("%d", team.TotalUsersCount),
			fmt.Sprintf("%d", team.TotalConferencesCount),
			fmt.Sprintf("%d", team.TotalIvrsCount),
			fmt.Sprintf("%d", team.TotalTagsCount),
			fmt.Sprintf("%d", team.TotalGroupsCount),
		})

		table.Render()

		return nil

	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
