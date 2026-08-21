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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves all IVR (Interactive Voice Response) configurations for your team.",
	Long: `Retrieves all IVR (Interactive Voice Response) configurations for your team.
An IVR defines an automated call flow — menu options, routing rules, welcome messages,
business hours, and queue behavior. Each IVR has one or more scenarios (call flow variants).
The total count is in the list_count field.

Permission:

	IVRs Read required.

Monitoring impact:
	
	OFF: Returns only IVRs assigned to you (IVRs using your numbers).
	ON: Returns all IVRs in the team.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		ivrs, reqInfo, err := httpClient.GetIVRs()
		if err != nil {
			return err
		}

		// Case no IVRs found
		if ivrs == nil {
			fmt.Println("No IVRs found")
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
			if err := formats.Table(os.Stdout, ivrs.List); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}
		default:
			if err := formats.JSON(os.Stdout, ivrs.List); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	ivrsCmd.AddCommand(listCmd)
	listCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
