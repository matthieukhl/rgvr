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

package numbers

import (
	"fmt"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// assignCmd represents the assign command
var assignCmd = &cobra.Command{
	Use:   "assign <number> (--user <user_id> | --ivr <ivr_id> | --conference <conference_id>)",
	Short: "Reassign a phone number to a different target: a user, an IVR, or a conference room.",
	Long: `Reassign a phone number to a different target: a user, an IVR, or a conference room.
Only one target field should be provided in the request body — user_id, ivr_id, or conference_id.
The number must already be provisioned in your team. Reassigning a number immediately changes call routing for that number.

Permission:

	Numbers Read required (assignment uses read-level permission).

Monitoring impact:

	OFF: Limited assignment capabilities — restricted to your own numbers.
	ON: Full assignment capabilities across the entire team.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.MarkFlagsMutuallyExclusive("user", "ivr", "conference")
		number := args[0]

		userID, _ := cmd.Flags().GetInt64("user")
		ivrID, _ := cmd.Flags().GetInt64("ivr")
		conferenceID, _ := cmd.Flags().GetInt64("conference")

		assignment := models.NumberAssignment{
			UserID:       userID,
			IvrID:        ivrID,
			ConferenceID: conferenceID,
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		reqInfo, err := httpClient.AssignNumber(number, &assignment)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		fmt.Printf("Number %s assigned successfully\n", number)
		return nil
	},
}

func init() {
	numbersCmd.AddCommand(assignCmd)
	assignCmd.Flags().Int64P("user", "u", 0, "Assign the number to a specific user by user ID.")
	assignCmd.Flags().Int64P("ivr", "i", 0, "Assign the number to a specific IVR by IVR ID.")
	assignCmd.Flags().Int64P("conference", "c", 0, "Assign the number to a specific conference room by conference ID.")
}
