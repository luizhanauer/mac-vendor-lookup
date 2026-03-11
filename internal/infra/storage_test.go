package infra

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/luizhanauer/mac-vendor-lookup/internal/domain"
)

func TestStorage_SaveShards(t *testing.T) {
	tempDir := t.TempDir()
	storage := &Storage{BaseDir: tempDir}

	shards := make(map[domain.OUI]*domain.ShardAggregate)
	oui, _ := domain.NewOUI("A8:23:FE")

	shards[oui] = &domain.ShardAggregate{
		BaseVendor: "Apple, Inc.",
		Blocks: []domain.VendorInfo{
			{
				MacPrefix:   "A823FE08",
				CompanyName: "Specific Apple Dept",
				MaskLength:  32,
			},
		},
	}

	err := storage.SaveShards(shards)
	if err != nil {
		t.Fatalf("unexpected error saving shards: %v", err)
	}

	expectedPath := filepath.Join(tempDir, "v1", "A8", "23", "FE.json")
	fileInfo, err := os.Stat(expectedPath)
	if err != nil {
		t.Fatalf("expected file to exist at %s: %v", expectedPath, err)
	}

	if fileInfo.IsDir() {
		t.Fatalf("expected %s to be a file, not a directory", expectedPath)
	}

	data, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("failed to read generated file: %v", err)
	}

	var result domain.ShardAggregate
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if result.BaseVendor != "Apple, Inc." {
		t.Errorf("expected base vendor 'Apple, Inc.', got '%s'", result.BaseVendor)
	}

	if len(result.Blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(result.Blocks))
	}

	if result.Blocks[0].CompanyName != "Specific Apple Dept" {
		t.Errorf("expected block company 'Specific Apple Dept', got '%s'", result.Blocks[0].CompanyName)
	}
}
