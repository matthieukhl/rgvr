// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tags

import (
	"fmt"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create --name <name> --color <color> [--description <text>]",
	Short: "Creates a new tag for your team.",
	Long: `
Creates a new tag for your team.
Tags are team-wide labels used to categorize calls.
Each tag has a name, an optional description, and a color chosenfrom a predefined palette.
Color options are: yellow, brown, pink, red, blue, green, grey, purple and orange.
Tag names should be unique within the team.
Once created, tags can be attached to calls for filtering and reporting.

Permission:

	IVRs Write required.

Monitoring:

	Not needed. Tag creation works regardless of the Monitoring flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Mark 'color' and 'name' flags as required
		cmd.MarkFlagsRequiredTogether("name", "color")

		name, _ := cmd.Flags().GetString("name")
		color, _ := cmd.Flags().GetString("color")
		description, _ := cmd.Flags().GetString("description")

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		// Validate name flag
		if name == "" {
			return fmt.Errorf("'name' flag cannot be empty")
		}

		// Validate color
		tagColor, err := models.ParseTagColor(color)
		if err != nil {
			return err
		}

		// Build tag struct
		tag := models.Tag{
			Name:        name,
			Color:       tagColor,
			Description: description,
		}

		// HTTP request
		reqInfo, err := httpClient.CreateTag(&tag)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return fmt.Errorf("checking verbose flag: %w", err)
		}

		return nil

	},
}

func init() {
	tagsCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "")
	createCmd.Flags().StringP("color", "c", "", "")
	createCmd.Flags().StringP("description", "d", "", "")
}
