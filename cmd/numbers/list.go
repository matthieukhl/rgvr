// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package numbers

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
	Short: "Retrieves the list of phone numbers provisioned for your team.",
	Long: `Retrieves the list of phone numbers provisioned for your team.
Each number includes its assignment (user, IVR, conference, fax, or unassigned), 
ountry, and configuration. The total count is in the list_count field.

Five boolean filters allow you to select numbers by assignment type. By default,
is_available is true and all others are false. Set a filter to true to include that category.
If is_available is false, at least one other filter must be true.

Permission:

	Numbers Read required.

Monitoring impact:

	OFF: Returns only your own numbers (numbers assigned to your user). All filter parameters are ignored
	— the response always contains your numbers regardless of filters.
	ON: Returns all team numbers with full filter support (is_ivr, is_user, is_conference, is_fax, is_available).`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		numbers, reqInfo, err := httpClient.GetNumbers()
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
			if err := formats.Table(os.Stdout, numbers.List); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}

		default:
			if err := formats.JSON(os.Stdout, numbers.List); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	numbersCmd.AddCommand(listCmd)
	listCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
