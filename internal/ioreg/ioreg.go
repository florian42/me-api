package ioreg

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/florian42/me-api/internal/cmd"
)

func GetIdleTime(runner cmd.CommandRunner) (time.Duration, error) {
	out, err := runner.Run("ioreg", "-c", "IOHIDSystem", "-l")
	if err != nil {
		return 0, err
	}

	re := regexp.MustCompile(`"HIDIdleTime"\s*=\s*(\d+)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		return 0, errors.New("HIDIdleTime not found")
	}
	nanoseconds, _ := strconv.ParseUint(matches[1], 10, 64)
	return time.Duration(nanoseconds), nil
}
