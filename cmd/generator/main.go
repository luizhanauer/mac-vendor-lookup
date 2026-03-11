package main

import (
	"flag"
	"log"
	"os"

	"github.com/luizhanauer/mac-vendor-lookup/internal/infra"
)

func main() {
	inputFile := flag.String("input", "manuf.txt", "Path to the Wireshark manuf file")
	outputDir := flag.String("output", "public", "Directory to output the static API")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	parser := &infra.Parser{}
	shards, err := parser.Parse(file)
	if err != nil {
		log.Fatalf("Error parsing manuf file: %v\n", err)
	}

	storage := &infra.Storage{BaseDir: *outputDir}
	if err := storage.SaveShards(shards); err != nil {
		log.Fatalf("Error saving shards: %v\n", err)
	}

	log.Printf("Successfully generated static API for %d OUIs in %s directory\n", len(shards), *outputDir)
}
