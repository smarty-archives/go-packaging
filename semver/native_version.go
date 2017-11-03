package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type NativeVersion struct {
	Major int
	Minor int
	Patch int
	dirty bool
}

func ParseNative(version string) (*NativeVersion, error) {
	components := strings.Split(version, "-")
	dirty := len(components) == 3
	parts := strings.Split(components[0], ".")

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

	info := NativeVersion{}
	info.dirty = dirty
	info.Major = parsed[0]
	info.Minor = parsed[1]
	if len(parsed) > 2 {
		info.Patch = parsed[2]
	}

	return &info, nil
}

func (this *NativeVersion) IncrementPatch() *NativeVersion {
	if !this.dirty {
		return this
	}

	return &NativeVersion{
		Major: this.Major,
		Minor: this.Minor,
		Patch: this.Patch + 1,
		dirty: false,
	}
}

func (this *NativeVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", this.Major, this.Minor, this.Patch)
}
