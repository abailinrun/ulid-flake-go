package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	ulidflake "github.com/abailinrun/ulid-flake-go/ulidflakescalable"
)

func main() {
	fmt.Println(`
    ██╗░░░██╗██╗░░░░░██╗██████╗░░░░░░░███████╗██╗░░░░░░█████╗░██╗░░██╗███████╗░░░░░░░██████╗
    ██║░░░██║██║░░░░░██║██╔══██╗░░░░░░██╔════╝██║░░░░░██╔══██╗██║░██╔╝██╔════╝░░░░░░██╔════╝
    ██║░░░██║██║░░░░░██║██║░░██║█████╗█████╗░░██║░░░░░███████║█████═╝░█████╗░░█████╗╚█████╗░
    ██║░░░██║██║░░░░░██║██║░░██║╚════╝██╔══╝░░██║░░░░░██╔══██║██╔═██╗░██╔══╝░░╚════╝░╚═══██╗
    ╚██████╔╝███████╗██║██████╔╝░░░░░░██║░░░░░███████╗██║░░██║██║░╚██╗███████╗░░░░░░██████╔╝
    ░╚═════╝░╚══════╝╚═╝╚═════╝░░░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝░░░░░░╚═════╝░
    `)

	// Define command-line flags
	generateFlag := flag.Bool("generate", false, "Generate a new Ulid-Flake")
	parseFlag := flag.String("parse", "", "Parse a Ulid-Flake string")
	epochFlag := flag.String("epoch", "2024-01-01T00:00:00Z", "Set the custom epoch time (e.g., 2024-01-01T00:00:00Z)")
	entropyFlag := flag.Int("entropy", 1, "Set the custom entropy size (default: 1)")
	sidFlag := flag.Int64("sid", 0, "Set the custom scalability ID (default: 0)")

	flag.Parse()

	// Set custom configuration if provided
	var opts []ulidflake.Option
	if *epochFlag != "" {
		epochTime, err := time.Parse(time.RFC3339, *epochFlag)
		if err != nil {
			log.Fatalf("Invalid epoch time format: %v", err)
		}
		opts = append(opts, ulidflake.WithEpochTime(epochTime))
	}
	if *entropyFlag != 0 {
		opts = append(opts, ulidflake.WithEntropySize(*entropyFlag))
	}
	if *sidFlag != 0 {
		opts = append(opts, ulidflake.WithSID(*sidFlag))
	}
	if err := ulidflake.SetConfig(opts...); err != nil {
		log.Fatalf("Failed to set config: %v", err)
	}

	// Generate a new Ulid-Flake
	if *generateFlag {
		ulid, err := ulidflake.New()
		if err != nil {
			log.Fatalf("Failed to generate Ulid-Flake: %v", err)
		}
		fmt.Printf("Generated Ulid-Flake:\n")
		fmt.Printf("  Base32:     %s\n", ulid.String())
		fmt.Printf("  Integer:    %d\n", ulid.Int())
		fmt.Printf("  Timestamp:  %d\n", ulid.Timestamp())
		fmt.Printf("  Randomness: %d\n", ulid.Randomness())
		fmt.Printf("  SID:        %d\n", ulid.SID())
		fmt.Printf("  Hex:        %s\n", ulid.Hex())
		fmt.Printf("  Bin:        %s\n", ulid.Bin())
		os.Exit(0)
	}

	// Parse an existing Ulid-Flake string
	if *parseFlag != "" {
		ulid, err := ulidflake.Parse(*parseFlag)
		if err != nil {
			log.Fatalf("Failed to parse Ulid-Flake: %v", err)
		}
		fmt.Printf("Parsed Ulid-Flake:\n")
		fmt.Printf("  Base32:     %s\n", ulid.String())
		fmt.Printf("  Integer:    %d\n", ulid.Int())
		fmt.Printf("  Timestamp:  %d\n", ulid.Timestamp())
		fmt.Printf("  Randomness: %d\n", ulid.Randomness())
		fmt.Printf("  SID:        %d\n", ulid.SID())
		fmt.Printf("  Hex:        %s\n", ulid.Hex())
		fmt.Printf("  Bin:        %s\n", ulid.Bin())
		os.Exit(0)
	}

	// Show usage if no flags are provided
	flag.Usage()
	os.Exit(1)
}
