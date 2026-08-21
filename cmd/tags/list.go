// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tags

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
	Short: "Retrieves all tags defined for your team.",
	Long: `Retrieves all tags defined for your team. Tags are labels that can be attached to calls
for categorization and reporting purposes (e.g., "support", "sales", "follow-up").
Tags are team-wide — all users share the same tag set. The total count is in the list_count field.

Permission:

	IVRs Read required.

Monitoring:

	Not needed. All team tags are returned regardless of the Monitoring flag.

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		tags, reqInfo, err := httpClient.GetTags()
		if err != nil {
			return err
		}

		if tags.ListCount == 0 {
			fmt.Println("No tags found.")
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
			if err := formats.Table(os.Stdout, tags.List); err != nil {
				return fmt.Errorf("printing table: %w", err)
			}
		default:
			if err := formats.JSON(os.Stdout, tags.List); err != nil {
				return fmt.Errorf("printing JSON: %w", err)
			}
		}

		return nil
	},
}

func init() {
	tagsCmd.AddCommand(listCmd)
	listCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
