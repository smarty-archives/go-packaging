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
	components := strings.Split(strings.TrimSpace(value), "-")
	if len(components) == 2 {
		return parseClean(components)
	} else if len(components) >= 4 {
		return parseDirty(components)
	} else {
		return malformed()
	}
}

func parseClean(components []string) (UpstreamVersion, error) {
	revision, err := strconv.Atoi(components[1])
	if err != nil {
		return malformed()
	}

	return UpstreamVersion{
		Version:  components[0],
		Revision: revision,
	}, nil
}

func parseDirty(components []string) (UpstreamVersion, error) {
	revision, err := strconv.Atoi(components[len(components)-3])
	if err != nil {
		return malformed()
	}

	return UpstreamVersion{
		Version:  strings.Join(components[0:len(components)-3], "-"),
		Revision: revision,
		dirty:    true,
	}, nil
}

func malformed() (UpstreamVersion, error) {
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
