// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package users

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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all team users",
	Long: `Retrieves users from your team. Each user object includes
the user's profile information, assigned phone numbers, status, and configuration.
The total number of users is indicated in the list_count field.`,
	RunE: func(cd *cobra.Command, args []string) error {
		path := "/users"

		client := cd.Context().Value(client.ClientContextKey).(*client.Client)

		resp, reqInfo, err := client.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

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

		if err = flags.IsVerbose(cd, reqInfo); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	usersCmd.AddCommand(listCmd)
}
