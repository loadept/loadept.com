package config

import (
	"sync"
	"time"
)

var (
	ApplicationUptime time.Time
	onceConfig        sync.Once
)

func LoadConfig() {
	onceConfig.Do(func() {
		ApplicationUptime = time.Now()
	})
}
