package main

import (
	"fmt"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Dirty bool
}

func (this Version) String() string {
	return fmt.Sprintf("%d.%d.%d", this.Major, this.Minor, this.Patch)
}

func (this Version) IncrementMajor() Version {
	return Version{Major: this.Major + 1}
}

func (this Version) IncrementMinor() Version {
	return Version{Major: this.Major, Minor: this.Minor + 1}
}

func (this Version) IncrementPatch() Version {
	return Version{Major: this.Major, Minor: this.Minor, Patch: this.Patch + 1}
}

func (this Version) Increment(how string) Version {
	switch strings.ToLower(how) {
	case "major":
		return this.IncrementMajor()
	case "minor":
		return this.IncrementMinor()
	default:
		return this.IncrementPatch()
	}
}
