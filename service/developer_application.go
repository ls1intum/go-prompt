package service

import (
	"context"
	"go-prompt/db"

	"github.com/dgraph-io/ristretto"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	q     *db.Queries
	cache *ristretto.Cache
}

func NewService(q *db.Queries) *Service {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	return &Service{q, cache}
}

func (s Service) GetAllDeveloperApplications(ctx context.Context) ([]db.ListDeveloperApplicationRow, error) {
	// Check whether the value is present in the cache
	value, found := s.cache.Get("all_developer_applications")
	if found {
		return value.([]db.ListDeveloperApplicationRow), nil
	}
	applications, err := s.q.ListDeveloperApplication(ctx)
	if err != nil {
		log.WithError(err).Error("Error getting all developer applications")
		return nil, err
	}
	s.cache.Set("all_developer_applications", applications, 0)
	return applications, nil
}
