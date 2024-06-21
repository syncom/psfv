package softball

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

var ageGroup string

var mountToPlateDistanceAndOffsetInFeet = map[string][2]float64{
	"8u":   {30.0, 3.0},
	"10u":  {35.0, 4.0},
	"12u":  {40.0, 5.0},
	"14u+": {43.0, 6.0},
}

var SoftballCmd = &cobra.Command{
	Use:   "softball",
	Short: "Analyze fastpitch softball pitching videos to estimate pitching speed",
	Long: `Analyze fastpitch softball pitching videos to estimate pitching speed
	
	Input video file should contain a softball pitcher throwing a pitch, where the
	clip starts when the ball leaves the pitcher's hand, and ends when it crosses
	the home plate. The video is analyzed to determine the speed of the pitch.
	
	Estimation is based on the age group of the game (8u, 10u, 12u, 14u and
	above)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vpath := args[0]
		duration := getDurationOfVideo(vpath)
		fmt.Printf("Duration of video: %f\n", duration)
		fmt.Printf("Age group: %s\n", strings.ToLower(ageGroup))
		speed := getAveragePitchingSpeed(ageGroup, duration)
		fmt.Printf("Average pitching speed: %f mph (%f kph)\n", speed, MphToKph(speed))
	},
}

func init() {
	SoftballCmd.Flags().StringVarP(&ageGroup, "age-group", "a", "14u+",
		"Age group of the game (8u, 10u, 12u, 14u+)")
}

// getDurationOfVideo returns the duration of the video in seconds
func getDurationOfVideo(path string) float64 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening test file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Fatalf("Error getting data: %v", err)
	}

	return data.Format.DurationSeconds
}

// getAveragePitchingSpeed returns the average pitching speed in mph
func getAveragePitchingSpeed(ageGroup string, flyingTime float64) float64 {
	ag := strings.ToLower(ageGroup)
	dist_offset, ok := mountToPlateDistanceAndOffsetInFeet[ag]
	if !ok {
		log.Fatalf("Invalid age group: %s", ageGroup)
	}
	mountToPlateDistance := dist_offset[0]
	offset := dist_offset[1]
	// Feet per second
	speed_fps := (mountToPlateDistance - offset) / flyingTime
	// Miles per hour. 1 hour = 3600 seconds, 1 mile = 5280 feet.
	speed_mph := speed_fps * 3600 / 5280

	return speed_mph
}

// MphToKph converts miles per hour to kilometers per hour
func MphToKph(mph float64) float64 {
	return mph * 1.60934
}
