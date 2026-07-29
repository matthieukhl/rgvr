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

package teams

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plans",
	Short: "Retrieve your team's plans information.",
	Long: `Retrieves a list of plans associated with your team,
including the number of licences used and the total number
of licences for each plan.`,
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/teams/plans_data"

		client := cd.Context().Value(internal.ClientContextKey).(*internal.Client)

		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving plan data: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var planData []models.PlanData
		if err := json.NewDecoder(resp.Body).Decode(&planData); err != nil {
			return fmt.Errorf("decoding plan data: %w", err)
		}

		format, err := cd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if err := formats.Check(format); err != nil {
			return err
		}

		switch format {
		case "table":
			if err := formats.Table(os.Stdout, planData); err != nil {
				return err
			}
		default:
			if err := formats.JSON(os.Stdout, planData); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	TeamsCmd.AddCommand(planCmd)
}
