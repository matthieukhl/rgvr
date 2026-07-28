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
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plans",
	Short: "Retrieve your team's plans information.",
	Long: `Retrieves a list of plans associated with your team,
including the number of licences used and the total number
of licences for each plan.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "/teams/plans_data"

		client := cmd.Context().Value(clientContextKey).(*internal.Client)

		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving plan data: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var planData models.PlanDataRaw
		if err := json.NewDecoder(resp.Body).Decode(&planData); err != nil {
			return fmt.Errorf("decoding plan data: %w", err)
		}

		table := tablewriter.NewTable(os.Stdout, tablewriter.WithHeaderAutoFormat(tw.Off))
		table.Header([]string{
			"Plan ID",
			"Plan Name",
			"Number of Licences Used",
			"Total Number of Licences",
		})

		for _, plan := range planData {
			table.Append([]string{
				fmt.Sprintf("%d", plan.PlanID),
				plan.PlanName,
				fmt.Sprintf("%d", plan.NbLicencesUsed),
				fmt.Sprintf("%d", plan.NbLicences),
			})
		}
		table.Render()
		return nil
	},
}

func init() {
	teamsCmd.AddCommand(planCmd)
}
