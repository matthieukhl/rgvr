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

package users

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invites one or more users to the team and creates all associated resources.",
	Long: `
Invites one or more users to the team and creates all associated resources (user accounts, phone numbers, voicemail, etc.).
Each invited user receives an activation email.
This operation consumes available licences from your subscription plan — ensure you have enough licences before calling this endpoint.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Not needed.

Warning — Sensitive operation:

	This route does not enforce any specific permission category.
	Any valid API key can invite new users to the team, consuming licences
	and generating activation emails.
	Consider restricting API key distribution to prevent unauthorized user creation.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch user ID from config to populate UserInvitationPayload.InvitedBy field
		userID := viper.GetInt("user_id")

		// UserInvite object creation
		var userInvites []models.UserInvite

		// Check if --file flag has been called by user
		filename, err := cmd.Flags().GetString("file")
		if err != nil {
			return fmt.Errorf("parsing file flags: %w", err)
		}

		// Check if there is a filename that has been input by the user
		if filename != "" {
			// Open file
			file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
			if err != nil {
				return fmt.Errorf("opening file: %w", err)
			}

			fileExtension := filepath.Ext(filename)

			switch fileExtension {
			case ".json":
				if userInvites, err = buildUserInvitePayloadFromJSON(file, userInvites); err != nil {
					return err
				}
			case ".csv":
				userInvites, err = buildUserInvitePayloadFromCSV(file)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("invalid file format. Supported file formats: csv/json")
			}
		} else {
			userInvites = buildUserInvitePayloadFromFlags(cmd)
		}

		if len(userInvites) == 0 {
			return fmt.Errorf("no invitations to send")
		}

		invitations := models.UserInvitationPayload{
			InvitedBy: userID,
			Users:     userInvites,
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := httpClient.InviteUser(&invitations)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		for _, user := range resp {
			fmt.Printf("Invitation sent to: %s (User ID: %d)\n", user.Email, user.ID)
		}

		return nil
	},
}

func init() {
	usersCmd.AddCommand(inviteCmd)

	inviteCmd.Flags().String("file", "", "")

	// Ad hoc flags
	inviteCmd.Flags().Int("number", 0, "")
	inviteCmd.Flags().String("email", "", "")
	inviteCmd.Flags().Int("plan", 0, "")

	// Flags constraints
	inviteCmd.MarkFlagsMutuallyExclusive("file", "number")
	inviteCmd.MarkFlagsMutuallyExclusive("file", "email")
	inviteCmd.MarkFlagsMutuallyExclusive("file", "plan")
	inviteCmd.MarkFlagsRequiredTogether("number", "email", "plan")

}

func buildUserInvitePayloadFromJSON(file *os.File, userInvites []models.UserInvite) ([]models.UserInvite, error) {
	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", file.Name(), err)
	}

	if err = json.Unmarshal(data, &userInvites); err != nil {
		return nil, fmt.Errorf("decoding data from file %s: %w", file.Name(), err)
	}

	return userInvites, nil
}

func buildUserInvitePayloadFromCSV(file *os.File) ([]models.UserInvite, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading .csv file: %w", err)
	}

	var userInvites []models.UserInvite

	for _, record := range records[1:] {
		number, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid number format: %w", err)
		}

		email := record[1]

		planId, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("invalid plan ID format: %w", err)
		}

		userInvite := models.UserInvite{
			Number: number,
			Email:  email,
			PlanID: planId,
		}

		userInvites = append(userInvites, userInvite)
	}

	return userInvites, nil
}

func buildUserInvitePayloadFromFlags(cmd *cobra.Command) []models.UserInvite {
	number, _ := cmd.Flags().GetInt("number")
	email, _ := cmd.Flags().GetString("email")
	planId, _ := cmd.Flags().GetInt("plan")

	userInvite := models.UserInvite{
		Number: number,
		Email:  email,
		PlanID: planId,
	}

	var userInvites []models.UserInvite

	userInvites = append(userInvites, userInvite)

	return userInvites
}
