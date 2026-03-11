package infra

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/luizhanauer/mac-vendor-lookup/internal/domain"
)

// Parser converts the manuf file into a map of ShardAggregates.
type Parser struct{}

// Parse reads the manuf file and groups entries by OUI.
func (p *Parser) Parse(r io.Reader) (map[domain.OUI]*domain.ShardAggregate, error) {
	shards := make(map[domain.OUI]*domain.ShardAggregate)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		p.processLine(line, shards)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed scanning file: %w", err)
	}

	return shards, nil
}

func (p *Parser) processLine(line string, shards map[domain.OUI]*domain.ShardAggregate) {
	if line == "" {
		return
	}
	if strings.HasPrefix(line, "#") {
		return
	}

	parts := strings.SplitN(line, "\t", 3)
	if len(parts) < 2 {
		return
	}

	rawPrefix := strings.TrimSpace(parts[0])
	companyName := strings.TrimSpace(parts[1])
	if len(parts) == 3 {
		companyName = strings.TrimSpace(parts[2]) // Use full name if available in comments
	}

	p.addToShards(rawPrefix, companyName, shards)
}

func (p *Parser) addToShards(rawPrefix, companyName string, shards map[domain.OUI]*domain.ShardAggregate) {
	oui, err := domain.NewOUI(rawPrefix)
	if err != nil {
		return
	}

	if _, exists := shards[oui]; !exists {
		shards[oui] = &domain.ShardAggregate{}
	}

	cleanPrefix, mask := p.extractPrefixAndMask(rawPrefix)

	if mask == 24 {
		shards[oui].BaseVendor = companyName
		return
	}

	shards[oui].Blocks = append(shards[oui].Blocks, domain.VendorInfo{
		MacPrefix:   cleanPrefix,
		CompanyName: companyName,
		MaskLength:  mask,
	})
}

func (p *Parser) extractPrefixAndMask(raw string) (string, int) {
	parts := strings.Split(raw, "/")
	clean := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(parts[0], ":", ""), "-", ""))

	if len(parts) > 1 {
		parsedMask, err := strconv.Atoi(parts[1])
		if err == nil {
			return clean, parsedMask
		}
	}

	if len(clean) > 6 {
		return clean, len(clean) * 4 // Fallback heuristic for unmasked hex longer than 6 chars
	}

	return clean, 24
}
