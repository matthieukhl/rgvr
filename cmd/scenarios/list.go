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

package scenarios

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves all scenarios across all IVRs in a single request.",
	Long: `Retrieves all scenarios across all IVRs in a single request.
This is a convenience alternative to calling 'rgvr ivrs scenarios list' for each IVR individually.
Useful for building a global overview of all call flows configured in the team.

Permission:

	IVRs Read required.

Monitoring impact:

	OFF: Returns only scenarios from IVRs assigned to you.
	ON: Returns all scenarios across all team IVRs.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	scenariosCmd.AddCommand(listCmd)
}
