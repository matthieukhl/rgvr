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

package auth

import (
	"fmt"

	"github.com/matthieukhl/rgvr/internal/config"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save your Ringover API key locally",
	Long: `Prompts for your Ringover API key and stores it in
~/.config/rgvr/config.yaml so you don't need to pass it on every command.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		configPath, err := cfg.Write()
		if err != nil {
			return err
		}

		fmt.Printf("API key saved to %s\n", configPath)
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
}
