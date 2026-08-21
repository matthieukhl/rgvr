// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package tags implements commands for managing
// Ringover tags via the rgvr CLI.
package tags

import (
	"context"
	"fmt"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage call tags.",
	Long: `Manage call tags. Tags are labels you can assign to calls for categorization and reporting.
Each tag has an ID, name, color, description, and creation date.

Permissions required:

	Read (IVRs R): List tags.
	Write (IVRs W): Create new tags.

Monitoring impact:

	None — tags are always team-wide.

`,
	PersistentPreRunE: func(cd *cobra.Command, args []string) error {
		httpClient, err := client.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		ctx := context.WithValue(cd.Context(), client.ClientContextKey, httpClient)
		cd.SetContext(ctx)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(tagsCmd)
}
