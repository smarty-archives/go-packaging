package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type UpstreamVersion struct {
	Version  string
	Revision int
	dirty    bool
}

func ParseUpstream(value string) (UpstreamVersion, error) {
	value = strings.TrimSpace(value)
	components := strings.Split(value, "-")
	dirty := len(components) != 2

	if len(components) < 2 {
		return UpstreamVersion{}, errors.New("Malformed version.")
	}

	if len(components) >= 4 {
		revision, err := strconv.Atoi(components[len(components)-3])
		if err != nil {
			return UpstreamVersion{}, errors.New("Malformed version.")
		}

		return UpstreamVersion{
			Version:  strings.Join(components[0:len(components)-3], "-"),
			Revision: revision,
			dirty:    true,
		}, nil
	}

	if !dirty {
		revision, err := strconv.Atoi(components[1])
		if err != nil {
			return UpstreamVersion{}, errors.New("Malformed version.")
		}

		return UpstreamVersion{
			Version:  components[0],
			Revision: revision,
		}, nil
	}

	return UpstreamVersion{}, errors.New("Malformed version.")
}

func (this UpstreamVersion) Increment() UpstreamVersion {
	if !this.dirty {
		return this
	}

	return UpstreamVersion{
		Version:  this.Version,
		Revision: this.Revision + 1,
	}
}

func (this UpstreamVersion) String() string {
	return fmt.Sprintf("%s-%d", this.Version, this.Revision)
}
