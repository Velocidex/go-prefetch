package prefetch

import (
	"time"
)

type PrefetchInfo struct {
	Executable string `json:"Executable"`
	// Path contains a kernel path to the executable.
	// For Windows Apps, it apparently contains the app name instead.
	Path          string      `json:"Path"`
	FileSize      uint32      `json:"FileSize"`
	Hash          string      `json:"Hash"`
	Version       string      `json:"Version"`
	LastRunTimes  []time.Time `json:"LastRunTimes"`
	FilesAccessed []string    `json:"FilesAccessed"`
	RunCount      uint32      `json:"RunCount"`
}
