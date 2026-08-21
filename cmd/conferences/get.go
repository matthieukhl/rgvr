// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conferences

import (
	"fmt"
	"os"
	"strconv"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <conference_id>",
	Short: "Retrieves detailed information about a specific conference room by its identifier.",
	Long: `
Retrieves detailed information about a specific conference room by its identifier,
including its assigned phone number, participant PIN code, admin PIN code, name, and configuration.

Permission:

	No specific permission required. A valid API key is sufficient.

Monitoring:

	Required. Conference details are a team-level supervisory feature. Returns 401 Unauthorized if Monitoring is OFF on your API key.

`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conferenceID := args[0]

		if conferenceID == "" {
			return fmt.Errorf("conference ID cannot be empty")
		}

		// Convert conference ID from string to int
		conferenceIDInt, err := strconv.Atoi(conferenceID)
		if err != nil {
			return fmt.Errorf("parsing conference ID: %w", err)
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		conference, reqInfo, err := httpClient.GetConference(conferenceIDInt)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.Conference{*conference}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, conference); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	conferencesCmd.AddCommand(getCmd)
	getCmd.Flags().String("format", "json", "Choose the output's format: table / json")
}
