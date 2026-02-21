package lsappinfo

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/florian42/me-api/internal/cmd"
)

func GetFrontmostAppName(runner cmd.CommandRunner) (string, error) {
	// First, get the ASN (Application Serial Number) of the frontmost app
	out, err := runner.Run("lsappinfo", "front")
	if err != nil {
		return "", fmt.Errorf("failed to get frontmost app: %w", err)
	}

	asn := strings.TrimSpace(string(out))
	if asn == "" {
		return "", errors.New("no frontmost app found")
	}

	// Now get the info for that ASN
	out, err = runner.Run("lsappinfo", "info", "-only", "name", asn)
	if err != nil {
		return "", fmt.Errorf("failed to get app info: %w", err)
	}

	// Parse the output like: "CFBundleDisplayName"="Safari"
	re := regexp.MustCompile(`"CFBundleDisplayName"="([^"]+)"`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		// Try alternate key names
		re = regexp.MustCompile(`"LSDisplayName"="([^"]+)"`)
		matches = re.FindStringSubmatch(string(out))
		if len(matches) < 2 {
			return "", errors.New("could not parse app name from lsappinfo output")
		}
	}

	return matches[1], nil
}
