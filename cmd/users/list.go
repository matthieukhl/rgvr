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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/matthieukhl/rgvr/internal"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/formats"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all team users",
	Long: `Retrieves users from your team. Each user object includes
the user's profile information, assigned phone numbers, status, and configuration.
The total number of users is indicated in the list_count field.`,
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/users"

		client := cd.Context().Value(internal.ClientContextKey).(*internal.Client)

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

		var usersResponse models.ListResponse[models.User]

		if err := json.NewDecoder(resp.Body).Decode(&usersResponse); err != nil {
			return err
		}

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
			formats.Table(os.Stdout, usersResponse.List)
		default:
			formats.JSON(os.Stdout, usersResponse)
		}

		if err = flags.IsVerbose(cd, resp, duration); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	UsersCmd.AddCommand(listCmd)
}
