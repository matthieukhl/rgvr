// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ivrs

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/spf13/cobra"
)

// listScenariosCmd represents the list command for scenarios
var listScenariosCmd = &cobra.Command{
	Use:   "list <ivr_id>",
	Short: "Retrieves all scenarios attached to a specific IVR.",
	Long: `Retrieves all scenarios attached to a specific IVR.
A scenario defines a complete call flow  menu options, routing, and what happens at each step.
An IVR typically has one default scenario but can have multiple variants (e.g. business hours vs. after-hours).
Total count is in list_count.

Permission:

	IVRs Read required.

Monitoring impact:

	OFF: returns only scenarios from IVRs assigned to you
	ON: returns all scenarios for the specified IVR
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ivrID := args[0]

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		scenarios, reqInfo, err := httpClient.GetIVRScenarios(ivrID)
		if err != nil {
			return err
		}

		// Case no scenarios found
		if scenarios == nil {
			fmt.Println("No scenarios found")
			return nil
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
			if err := formats.Table(os.Stdout, scenarios.List); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}
		default:
			if err := formats.JSON(os.Stdout, scenarios); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	scenariosCmd.AddCommand(listScenariosCmd)
}
