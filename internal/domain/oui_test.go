package domain

import (
	"testing"
)

func TestNewOUI_Valid(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		expected OUI
	}{
		{"Standard colon separated", "00:50:C2", "0050C2"},
		{"Standard dash separated", "00-50-C2", "0050C2"},
		{"Lower case", "a8:23:fe", "A823FE"},
		{"Full MAC address", "a8:23:fe:08:37:ba", "A823FE"},
		{"Full MAC with mask", "00:50:C2:4B:40/36", "0050C2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oui, err := NewOUI(tt.raw)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if oui != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, oui)
			}
		})
	}
}

func TestNewOUI_Invalid(t *testing.T) {
	raw := "12:34" // Too short
	_, err := NewOUI(raw)

	if err == nil {
		t.Fatalf("expected error for invalid OUI length, got nil")
	}
}

func TestOUIShardPath(t *testing.T) {
	oui := OUI("A823FE")
	expectedDir := "A8/23"
	expectedFile := "FE.json"

	dir, file := oui.ShardPath()

	if dir != expectedDir {
		t.Errorf("expected dir %s, got %s", expectedDir, dir)
	}

	if file != expectedFile {
		t.Errorf("expected file %s, got %s", expectedFile, file)
	}
}
