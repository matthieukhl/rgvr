// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conferences

import (
	"fmt"
	"strconv"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/spf13/cobra"
)

// setPinCmd represents the setPin command
var setPinCmd = &cobra.Command{
	Use:   "set-pin <conference_id> <pincode>",
	Short: "Updates the participant PIN code for a specific conference room.",
	Long: `
Updates the participant PIN code for a specific conference room.
The new PIN code value is passed directly in the URL path as pincodeId.
Participants will need to use the new PIN to join the conference.
If the new PIN is the same as the current one, the request returns 304 Not Modified.

Permission:
	
	No specific permission required. A valid API key is sufficient.

Monitoring:

	Required. Conference configuration changes are team-level supervisory operations. Returns 401 Unauthorized if Monitoring is OFF on your API key.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		conferenceID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("please enter a valid conference ID")
		}

		pincode, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("pincode value must be an integer between 1 and 9999999999999999")
		}

		// Check pincode validity
		if pincode < 1 && pincode > 9999999999999999 {
			return fmt.Errorf("pincode value must be an integer between 1 and 9999999999999999")
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		reqInfo, err := httpClient.SetConferencePin(conferenceID, pincode)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		fmt.Println("Conference user PIN code successfully updated!")

		return nil

	},
}

func init() {
	conferencesCmd.AddCommand(setPinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setPinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setPinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
