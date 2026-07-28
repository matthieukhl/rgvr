/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plans",
	Short: "Retrieve your team's plans information.",
	Long: `Retrieves a list of plans associated with your team,
including the number of licences used and the total number
of licences for each plan.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "/teams/plans_data"

		client := cmd.Context().Value(clientContextKey).(*internal.Client)

		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving plan data: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var planData models.PlanDataRaw
		if err := json.NewDecoder(resp.Body).Decode(&planData); err != nil {
			return fmt.Errorf("decoding plan data: %w", err)
		}

		table := tablewriter.NewTable(os.Stdout, tablewriter.WithHeaderAutoFormat(tw.Off))
		table.Header([]string{
			"Plan ID",
			"Plan Name",
			"Number of Licences Used",
			"Total Number of Licences",
		})

		for _, plan := range planData {
			table.Append([]string{
				fmt.Sprintf("%d", plan.PlanID),
				plan.PlanName,
				fmt.Sprintf("%d", plan.NbLicencesUsed),
				fmt.Sprintf("%d", plan.NbLicences),
			})
		}
		table.Render()
		return nil
	},
}

func init() {
	teamsCmd.AddCommand(planCmd)
}
