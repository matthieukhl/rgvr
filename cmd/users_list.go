/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah

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

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/matthieukhl/rgvr/internal"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "/users"

		client := cmd.Context().Value(clientContextKey).(*internal.Client)

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

		// table := tablewriter.NewTable(os.Stdout, tablewriter.WithHeaderAutoFormat(tw.Off))
		// table.Header([]string{
		// 	"ID",
		// 	"First Name",
		// 	"Last Name",
		// 	"Email",
		// 	"Company",
		// 	"Picture",
		// 	"Numbers",
		// })

		// var numbers []int64

		// for _, user := range usersResponse.List {
		// 	for _, number := range user.Numbers {
		// 		numbers = append(numbers, number.Number)
		// 	}
		// }

		// for _, user := range usersResponse.List {
		// 	table.Append([]string{
		// 		fmt.Sprintf("%d", user.UserID),
		// 		user.Firstname,
		// 		user.Lastname,
		// 		user.Email,
		// 		user.Company,
		// 		user.Picture,
		// 		fmt.Sprintf("%v\n", numbers),
		// 	})
		// }

		// table.Render()

		if err := json.NewEncoder(os.Stdout).Encode(usersResponse); err != nil {
			return err
		}

		verbose, err := cmd.Flags().GetBool("verbose")
		if verbose {
			fmt.Fprintf(os.Stderr, "\nURL called: %s\n", resp.Request.URL.String())
			fmt.Fprintf(os.Stderr, "Query duration: %d ms\n", duration.Milliseconds())
		}

		return nil
	},
}

func init() {
	usersCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Display detailed information about each user")
}
