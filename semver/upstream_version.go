package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type UpstreamVersion struct {
	Base     string
	Revision int
}

func NewUpstreamVersion(base string, revision int) UpstreamVersion {
	return UpstreamVersion{Base: base, Revision: revision}
}

func ParseUpstream(base, raw string) (UpstreamVersion, error) {
	base = strings.TrimSpace(base)
	raw = strings.TrimSpace(raw)

	if len(base) == 0 {
		return malformed()
	} else if len(raw) == 0 {
		return NewUpstreamVersion(base, 0), nil
	} else if !strings.HasPrefix(raw, base) {
		return NewUpstreamVersion(base, 0), nil
	} else {
		return parseUpstream(base, raw[len(base):])
	}
}

func malformed() (UpstreamVersion, error) {
	return UpstreamVersion{}, errors.New("malformed version")
}

func parseUpstream(base, suffix string) (UpstreamVersion, error) {
	split := strings.Split(suffix, "-")
	if len(split) == 1 {
		return NewUpstreamVersion(base, 0), nil
	} else if len(split[0]) > 0 {
		return NewUpstreamVersion(base, 0), nil
	} else if len(split) == 2 {
		return increment(base, split[1], 0)
	} else if len(split) > 2 {
		return increment(base, split[1], 1)
	} else {
		return malformed()
	}
}

func increment(base, revision string, offset int) (UpstreamVersion, error) {
	if revision, err := strconv.Atoi(revision); err != nil {
		return malformed()
	} else {
		return NewUpstreamVersion(base, revision+offset), nil
	}
}

func (this UpstreamVersion) String() string {
	return fmt.Sprintf("%s-%d", this.Base, this.Revision)
}
