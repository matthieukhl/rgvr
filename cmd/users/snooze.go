// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"

	"github.com/spf13/cobra"
)

// snoozeCmd represents the snooze command
var snoozeCmd = &cobra.Command{
	Use:   "snooze",
	Short: "Manage users' snooze status",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snooze called")
	},
}

func init() {
	usersCmd.AddCommand(snoozeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snoozeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snoozeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
