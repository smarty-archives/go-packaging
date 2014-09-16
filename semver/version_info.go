package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type VersionInfo struct {
	Major int
	Minor int
	Patch int
}

func ParseVersion(version string) (*VersionInfo, error) {
	parts := strings.Split(version, ".")
	if len(parts) > 3 {
		return nil, errors.New("Malformed version, too many parts.")
	} else if len(parts) < 2 {
		return nil, errors.New("Malformed version, too few parts.")
	}

	parsed := []int{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if value, err := strconv.Atoi(part); err != nil {
			return nil, errors.New("Malformed version, all parts must be numeric.")
		} else {
			parsed = append(parsed, value)
		}
	}

	info := VersionInfo{}
	info.Major = parsed[0]
	info.Minor = parsed[1]
	if len(parsed) > 2 {
		info.Patch = parsed[2]
	}

	return &info, nil
}

func (this *VersionInfo) String() string {
	return fmt.Sprintf("%d.%d.%d", this.Major, this.Minor, this.Patch)
}
