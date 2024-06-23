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
	Short: "Pitch Speed From Video",
	Long: `Estimate the speed of a pitch from a video clip of a pitcher throwing a softball or baseball.
	
This tool depends on the ffprobe tool to extract video metadata. The video
clip should start when the ball leaves the pitcher's hand and end when it
crosses the home plate. The tool will estimate the speed of the pitch based on
the duration of the video clip and the age group of the game.`,
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
