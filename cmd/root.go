package psfv

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/syncom/psfv/cmd/softball"
	"github.com/syncom/psfv/version"
)

var rootCmd = &cobra.Command{
	Use:   "psfv",
	Short: "",
	Long:  ``,
	//Run: func(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("CLI error: %s", err)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(softball.SoftballCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of psfv",
	Long:  `Print the version of Pitching Speed From Video (psfv)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(("Pitching Speed From Video (psfv) v%s\n"), version.Version)
	},
}
