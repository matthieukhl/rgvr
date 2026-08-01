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

package teams

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// TeamsCmd represents the teams command
var TeamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Retrieve team information.",
	Long:  `Retrieves a complete team object containing lists of numbers, users, ivrs, conferences, tags and groups.`,
	PersistentPreRunE: func(cd *cobra.Command, args []string) error {
		httpClient, err := client.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		ctx := context.WithValue(cd.Context(), client.ClientContextKey, httpClient)
		cd.SetContext(ctx)
		return nil
	},
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/teams"
		client := cd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, err := client.Get(path)
		if err != nil {
			return fmt.Errorf("retrieving team information: %w", err)
		}
		defer resp.Body.Close()

		var team models.Team
		if err = json.NewDecoder(resp.Body).Decode(&team); err != nil {
			return fmt.Errorf("decoding team information: %w", err)
		}

		// This is ugly but necessary to be able to pass the `team` object
		// to formats.Table() func below
		var teamList = []models.Team{team}

		format, err := cd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if err = formats.Check(format); err != nil {
			return err
		}

		// Output result to stdout
		switch format {
		case "table":
			if err := formats.Table(os.Stdout, teamList); err != nil {
				return err
			}
		default:
			if err := formats.JSON(os.Stdout, team); err != nil {
				return err
			}
		}

		return nil

	},
}

func init() {
	cmd.RootCmd.AddCommand(TeamsCmd)

	TeamsCmd.PersistentFlags().String("format", "json", "Choose the output's format: table / json")
}
