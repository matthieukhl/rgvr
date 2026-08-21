// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/spf13/cobra"
)

// planningsCmd represents the plannings command
var planningsCmd = &cobra.Command{
	Use:   "plannings",
	Short: "Manage users' plannings",
	Long:  `TODO`,
}

func init() {
	usersCmd.AddCommand(planningsCmd)
}
