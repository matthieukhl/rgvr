// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package teams

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plans",
	Short: "Retrieve your team's plans information.",
	Long: `Retrieves a list of plans associated with your team,
including the number of licences used and the total number
of licences for each plan.`,
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/teams/plans_data"

		client := cd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving plan data: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var planData []models.PlanData
		if err := json.NewDecoder(resp.Body).Decode(&planData); err != nil {
			return fmt.Errorf("decoding plan data: %w", err)
		}

		format, err := cd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if err := formats.Check(format); err != nil {
			return err
		}

		switch format {
		case "table":
			if err := formats.Table(os.Stdout, planData); err != nil {
				return err
			}
		default:
			if err := formats.JSON(os.Stdout, planData); err != nil {
				return err
			}
		}

		if err := flags.IsVerbose(cd, reqInfo); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	TeamsCmd.AddCommand(planCmd)
}
