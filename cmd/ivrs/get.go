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
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <ivr_id>",
	Short: "Retrieves detailed information about a specific IVR by its identifier.",
	Long: `Retrieves detailed information about a specific IVR by its identifier,
including its configuration, assigned numbers, scenarios, welcome message settings,
business hours, and queue parameters. Use GET /ivrs to discover available IVR IDs.

Permission:

	IVRs Read required.

Monitoring impact:

	OFF: Returns the IVR only if it is assigned to you. Returns 404 otherwise.
	ON: Returns any IVR in the team.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ivrID := args[0]

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		ivr, reqInfo, err := httpClient.GetIVR(ivrID)
		if err != nil {
			return err
		}

		// Sanity check
		if ivr == nil {
			return fmt.Errorf("no IVR found")
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
			if err := formats.Table(os.Stdout, []models.IVR{*ivr}); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}

		default:
			if err := formats.JSON(os.Stdout, []models.IVR{*ivr}); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	ivrsCmd.AddCommand(getCmd)
	getCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
