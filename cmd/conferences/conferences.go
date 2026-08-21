// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package conferences implements commands for managing
// Ringover conference rooms via the rgvr CLI.
package conferences

import (
	"context"
	"fmt"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/spf13/cobra"
)

// conferencesCmd represents the conferences command
var conferencesCmd = &cobra.Command{
	Use:   "conferences",
	Short: "Manage conference rooms.",
	Long: `
Manage conference rooms.
List conferences and update access PIN codes.

Permissions required:

	No specific permission category, but Monitoring must be ON for all conference routes. Returns 401 if Monitoring is OFF.
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
	cmd.RootCmd.AddCommand(conferencesCmd)
}
