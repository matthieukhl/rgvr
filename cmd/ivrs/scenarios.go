package ivrs

import (
	"github.com/spf13/cobra"
)

// scenariosCmd represents the scenarios command
var scenariosCmd = &cobra.Command{
	Use:   "scenarios",
	Short: "Manage scenarios for IVRs.",
	Long:  `Manage scenarios for IVRs (Interactive Voice Response).`,
}

func init() {
	ivrsCmd.AddCommand(scenariosCmd)
}
