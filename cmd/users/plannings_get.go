/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package users

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// planningsGetCmd represents the planningsGet command
var planningsGetCmd = &cobra.Command{
	Use:   "get <user_id>",
	Short: "Retrieves a user's weekly availability planning — the time slots during which the user is reachable via Ringover.",
	Long: `Retrieves a user's weekly availability planning — the time slots during which the user is reachable via Ringover.
	The response contains a planning for each day of the week with start/end time ranges.
	If planning is not enabled in the user's settings, the is_planning field will be false and the time ranges will be empty.

Permission: 
	
	'Users Read' required.

Monitoring impact:

	OFF: Returns planning only if userId matches your own user ID. Requesting another user's planning returns 404.
	ON: Returns any user's planning in the team.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]

		path := fmt.Sprintf("/users/%s/plannings", userID)

		client := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		start := time.Now()
		resp, err := client.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		duration := time.Since(start)

		if resp.StatusCode != 200 {
			return fmt.Errorf("unexpected response from API: %s", resp.Status)
		}

		var planning models.Planning
		if err := json.NewDecoder(resp.Body).Decode(&planning); err != nil {
			return fmt.Errorf("decoding planning information: %w", err)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("retrieving format flag: %w", err)
		}

		switch format {
		case "table":
			if err = formats.Table(os.Stdout, []models.Planning{planning}); err != nil {
				return err
			}
		default:
			if err = formats.JSON(os.Stdout, planning); err != nil {
				return err
			}
		}

		if err := flags.IsVerbose(cmd, resp, duration); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	planningsCmd.AddCommand(planningsGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planningsGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planningsGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
