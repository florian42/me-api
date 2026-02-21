package presence

import (
	"time"

	"github.com/florian42/me-api/internal/cmd"
	"github.com/florian42/me-api/internal/ioreg"
	"github.com/florian42/me-api/internal/lsappinfo"
)

type PresenceStatus string

const (
	StatusActive PresenceStatus = "active"
	StatusIdle   PresenceStatus = "idle"
	StatusAway   PresenceStatus = "away"
	StatusLocked PresenceStatus = "locked"
	Unknown      PresenceStatus = "unknown"
)

func GetStatus(runner cmd.CommandRunner) PresenceStatus {
	status, err := isActive(runner)
	if err != nil {
		return "unknown"
	}
	if status {
		return "active"
	}

	return "idle"

}

const idleThreshold = 5 * time.Minute

func isActive(runner cmd.CommandRunner) (bool, error) {
	idleTime, err := ioreg.GetIdleTime(runner)
	if err != nil {
		return false, err
	}
	return idleTime < idleThreshold, nil
}

func GetFocusedApp(runner cmd.CommandRunner) (string, error) {
	return lsappinfo.GetFrontmostAppName(runner)
}
