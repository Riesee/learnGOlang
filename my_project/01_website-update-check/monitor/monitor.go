package monitor

import (
	"context"
	"time"
	"website-checker/fetch"
	"website-checker/hash"
	"website-checker/logger"
)

type Monitor struct {
	url string
	selector string
	interval time.Duration
	lastHash string
}

func (m *Monitor) Start(ctx context.Context) {
	ticker := time.NewTicker(m.interval)
	for {
	select {
	case <-ticker.C:
		// doc, err := fetch.Get(m.url, m.selector)
		// if err != nil {
		// 	logger.Info("Error fetching document", err)
		// 	continue
		// }

		

		// hash := hash.Hash(doc)

		// if hash != m.lastHash {
		// 	m.lastHash = hash
		// 	logger.Info("Hash changed:", hash)
		// }

		doc, err := fetch.GetWithBrowser(m.url, m.selector)
		if err != nil {
		logger.Error("Error fetching document", "error", err)
			continue
		}

		hash := hash.Hash(doc)

		if hash != m.lastHash {
			m.lastHash = hash
			// logger.Info("Hash changed", "hash", hash)
		}
	case <-ctx.Done():
		ticker.Stop()
		return
	}
	}
}



func New(url string, interval time.Duration, selector string) *Monitor {
	return &Monitor{
		url: url,
		interval: interval,
		selector: selector,
	}
}