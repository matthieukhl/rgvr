// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scenarios

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves all scenarios across all IVRs in a single request.",
	Long: `Retrieves all scenarios across all IVRs in a single request.
This is a convenience alternative to calling 'rgvr ivrs scenarios list' for each IVR individually.
Useful for building a global overview of all call flows configured in the team.

Permission:

	IVRs Read required.

Monitoring impact:

	OFF: Returns only scenarios from IVRs assigned to you.
	ON: Returns all scenarios across all team IVRs.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		scenarios, reqInfo, err := httpClient.GetScenarios()
		if err != nil {
			return err
		}

		// Case no scenarios found
		// The assertion scenarios.ListCount == 0 is my not so elegant
		// way of managing no scenarios returned by the API.
		if scenarios == nil || scenarios.ListCount == 0 {
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
	scenariosCmd.AddCommand(listCmd)
}
