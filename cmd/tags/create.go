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

package tags

import (
	"slices"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create --name <name> --color <hex> [--description <text>]",
	Short: "Creates a new tag for your team.",
	Long: `Creates a new tag for your team.
Tags are team-wide labels used to categorize calls.
Each tag has a name, an optional description, and a color chosen
from a predefined palette (18 hex colors).
Tag names should be unique within the team.
Once created, tags can be attached to calls for filtering and reporting.

Permission:

	IVRs Write required.

Monitoring:

	Not needed. Tag creation works regardless of the Monitoring flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		color, _ := cmd.Flags().GetString("color")
		description, _ := cmd.Flags().GetString("description")

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		contains := slices.Contains(colors, color)
		_, ok := colorFamilies[color]

		tag := models.Tag{
			Name:        name,
			Color:       models.TagColor(color),
			Description: description,
		}

		reqInfo, err := httpClient.CreateTag(tag)
	},
}

var colors = []string{
	"purple",
	"red",
	"green",
	"yellow",
	"grey",
	"brown",
	"blue",
	"orange",
	"pink",
}

var colorFamilies = map[string][]models.TagColor{
	"purple": {models.ColorPurple1, models.ColorPurple2, models.ColorPurple3},
	"red":    {models.ColorRed},
	"green":  {models.ColorGreen1, models.ColorGreen2, models.ColorGreen3, models.ColorGreen4},
	"yellow": {models.ColorYellow},
	"grey":   {models.ColorGrey1, models.ColorGrey2},
	"brown":  {models.ColorBrown},
	"blue":   {models.ColorBlue1, models.ColorBlue2, models.ColorBlue3},
	"orange": {models.ColorOrange1, models.ColorOrange2},
	"pink":   {models.ColorPink},
}

func init() {
	tagsCmd.AddCommand(createCmd)
}
