package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"theperiscope.org/avwx/internal/pkg/metars"
	"theperiscope.org/avwx/internal/pkg/tafs"
)

func main() {

	if len(os.Args) < 3 {
		showUsage()
		os.Exit(1)
	}

	cmd := strings.ToLower(os.Args[1])
	stations := os.Args[2:]

	var data []string
	var err error

	switch cmd {
	case "metar":
		data, err = metars.GetData(stations)
	case "taf":
		data, err = tafs.GetData(stations)
	default:
		log.Fatal("Invalid command.")
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(strings.Join(data, "\n"))
}

func showUsage() {
	fmt.Println("AVWX - simple tool to access METAR and TAF data")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("   avwx <metar|taf> <station1> [<...stationN>]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("   - single ICAO station:")
	fmt.Println("     avwx metar KORD")
	fmt.Println("     avwx taf KORD")
	fmt.Println()
	fmt.Println("   - multiple ICAO stations:")
	fmt.Println("     avwx metar KORD PHOG")
	fmt.Println("     avwx taf KORD PHOG")
	fmt.Println()
	fmt.Println("   - partial ICAO station name:")
	fmt.Println("     avwx metar PH*")
	fmt.Println("     avwx taf PH*")
	fmt.Println()
	fmt.Println("   - state/province: (use two-letter U.S. state or two-letter Canadian province)")
	fmt.Println("     ")
	fmt.Println("     avwx metar @il")
	fmt.Println("     avwx taf @il")
	fmt.Println()
	fmt.Println("   - country: (use two-letter country abbreviation)")
	fmt.Println("     avwx metar ~au")
	fmt.Println("     avwx taf ~au")
	fmt.Println()
	fmt.Println("   - mix and match:")
	fmt.Println("     avwx metar KORD CY* ~au @hi")
	fmt.Println("     taf KORD CY* ~au @hi")
}
