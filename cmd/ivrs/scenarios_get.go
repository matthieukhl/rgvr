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

package ivrs

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// getScenarioCmd represents the get command
var getScenarioCmd = &cobra.Command{
	Use:   "get <ivr_id> <scenario_id>",
	Short: "Retrieves detailed information about a specific scenario within an IVR.",
	Long: `Retrieves detailed information about a specific scenario within an IVR.
The response includes the full call flow definition — menu steps, routing rules,
audio prompts, and queue configuration. Both ivrId and scenarioId must match for
the scenario to be returned.

Permission:

	IVRs Read required.

Monitoring impact:

	OFF: Returns the scenario only if the parent IVR is assigned to you. Returns 404 otherwise.
	ON: Returns any scenario for any IVR in the team.
`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ivrID := args[0]
		scenarioID := args[1]

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		scenario, reqInfo, err := httpClient.GetIVRScenario(ivrID, scenarioID)
		if err != nil {
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		switch format {
		case "table":
			if err := formats.Table(os.Stdout, []models.Scenario{*scenario}); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}

		default:
			if err := formats.JSON(os.Stdout, []models.Scenario{*scenario}); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	scenariosCmd.AddCommand(getScenarioCmd)
}
