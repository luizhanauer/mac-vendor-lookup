package infra

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/luizhanauer/mac-vendor-lookup/internal/domain"
)

// Storage handles the persistence of shards into static JSON files.
type Storage struct {
	BaseDir string
}

// SaveShards writes the map of shards to the filesystem.
func (s *Storage) SaveShards(shards map[domain.OUI]*domain.ShardAggregate) error {
	for oui, aggregate := range shards {
		if err := s.saveSingleShard(oui, aggregate); err != nil {
			return fmt.Errorf("failed to save shard %s: %w", string(oui), err)
		}
	}
	return nil
}

func (s *Storage) saveSingleShard(oui domain.OUI, aggregate *domain.ShardAggregate) error {
	dir, file := oui.ShardPath()
	fullDir := filepath.Join(s.BaseDir, "v1", dir)

	if err := os.MkdirAll(fullDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", fullDir, err)
	}

	targetPath := filepath.Join(fullDir, file)
	data, err := json.Marshal(aggregate)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON for %s: %w", targetPath, err)
	}

	if err := os.WriteFile(targetPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", targetPath, err)
	}

	return nil
}
