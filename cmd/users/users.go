// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package users

import (
	"context"
	"fmt"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/spf13/cobra"
)

// UsersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage your Ringover team members",
	Long: `Manage your Ringover team members.
Retrieve user information, manage plannings, check availability (presences/snooze), and organize user blacklists.

Permissions required:

	Read (Users R): Get user details, plannings, presences, user blacklists.
	Write (Users W): Modify plannings, manage group members, manage user blacklists, modify snooze status.

Monitoring impact:

	OFF: GET requests return only your own user data, planning, and presence. Group listing returns only groups you belong to.
	ON: GET requests return all team users, plannings, and presences. You can modify any user's planning.
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
	cmd.RootCmd.AddCommand(usersCmd)
	usersCmd.PersistentFlags().String("format", "json", "Choose the output's format: table / json")
}
