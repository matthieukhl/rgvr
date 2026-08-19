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
