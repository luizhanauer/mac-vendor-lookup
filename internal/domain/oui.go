package domain

import (
	"fmt"
	"strings"
)

// OUI represents the first 3 bytes (24 bits) of a MAC address.
type OUI string

// NewOUI creates a validated OUI from a raw MAC string.
func NewOUI(raw string) (OUI, error) {
	clean := strings.ReplaceAll(raw, ":", "")
	clean = strings.ReplaceAll(clean, "-", "")
	clean = strings.ToUpper(clean)

	if len(clean) < 6 {
		return "", fmt.Errorf("invalid OUI length: %s", raw)
	}

	return OUI(clean[:6]), nil
}

// ShardPath returns the directory and filename for this OUI.
func (o OUI) ShardPath() (string, string) {
	s := string(o)
	return fmt.Sprintf("%s/%s", s[0:2], s[2:4]), fmt.Sprintf("%s.json", s[4:6])
}
