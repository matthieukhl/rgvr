// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package numbers

import (
	"context"
	"fmt"

	"github.com/matthieukhl/rgvr/cmd"
	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/spf13/cobra"
)

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "Manage your team's numbers.",
	Long: `List numbers, retrieve a number's details and assign 
a phone number to a user, an IVR or a conference.`,
	PersistentPreRunE: func(cd *cobra.Command, args []string) error {
		httpClient, err := client.NewClient()
		if err != nil {
			return fmt.Errorf("creating client: %w", err)
		}

		ctx := context.WithValue(cd.Context(), client.ClientContextKey, httpClient)
		cd.SetContext(ctx)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(numbersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
