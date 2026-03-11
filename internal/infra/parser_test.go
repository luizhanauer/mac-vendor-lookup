package infra

import (
	"strings"
	"testing"

	"github.com/luizhanauer/mac-vendor-lookup/internal/domain"
)

func TestParser(t *testing.T) {
	mockManuf := `
# Este é um comentário que deve ser ignorado
00:00:00	Xerox	Xerox Corporation
00:50:C2:4B:40/36	OUI_Specific	Specific Company Ltd.
`

	parser := &Parser{}
	reader := strings.NewReader(mockManuf)

	shards, err := parser.Parse(reader)
	if err != nil {
		t.Fatalf("unexpected error parsing manuf data: %v", err)
	}

	validateBaseVendor(t, shards, "00:00:00", "Xerox Corporation")
	validateIABBlock(t, shards, "00:50:C2:4B:40/36", "Specific Company Ltd.", 36)
}

func validateBaseVendor(t *testing.T, shards map[domain.OUI]*domain.ShardAggregate, rawMac, expectedVendor string) {
	oui, err := domain.NewOUI(rawMac)
	if err != nil {
		t.Fatalf("failed to create OUI for %s: %v", rawMac, err)
	}

	shard, exists := shards[oui]
	if !exists {
		t.Fatalf("expected shard for OUI %s to exist", oui)
	}

	if shard.BaseVendor != expectedVendor {
		t.Errorf("expected BaseVendor to be '%s', got '%s'", expectedVendor, shard.BaseVendor)
	}
}

func validateIABBlock(t *testing.T, shards map[domain.OUI]*domain.ShardAggregate, rawMac, expectedVendor string, expectedMask int) {
	oui, err := domain.NewOUI(rawMac)
	if err != nil {
		t.Fatalf("failed to create OUI for %s: %v", rawMac, err)
	}

	shard, exists := shards[oui]
	if !exists {
		t.Fatalf("expected shard for OUI %s to exist", oui)
	}

	if len(shard.Blocks) != 1 {
		t.Fatalf("expected 1 block in shard %s, got %d", oui, len(shard.Blocks))
	}

	block := shard.Blocks[0]
	if block.CompanyName != expectedVendor {
		t.Errorf("expected block company to be '%s', got '%s'", expectedVendor, block.CompanyName)
	}

	if block.MaskLength != expectedMask {
		t.Errorf("expected mask length to be %d, got %d", expectedMask, block.MaskLength)
	}
}
