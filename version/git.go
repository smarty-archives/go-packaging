package main

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

type Git struct{}

func (this *Git) CurrentVersion() (Version, error) {
	output, err := exec.Command("git", "describe").CombinedOutput()
	if err != nil {
		return parseGitDescribe(string(output) + " " + string(err.Error()))
	}
	return parseGitDescribe(string(output))
}

func parseGitDescribe(raw string) (version Version, err error) {
	raw = strings.TrimSpace(raw)
	if strings.HasPrefix(raw, "fatal: No names found, cannot describe anything.") {
		version.Dirty = true
		return version, nil
	}
	fields := strings.Split(raw, "-")
	version.Dirty = len(fields) > 1

	versionFields := strings.Split(fields[0], ".")
	if len(versionFields) < 3 {
		return Version{}, errors.New("At least 3 version fields are required (major.minor.patch).")
	}
	version.Major, err = strconv.Atoi(versionFields[0])
	if err != nil {
		return version, err
	}
	version.Minor, err = strconv.Atoi(versionFields[1])
	if err != nil {
		return version, err
	}
	version.Patch, err = strconv.Atoi(versionFields[2])
	if err != nil {
		return version, err
	}
	return version, nil
}

func (this *Git) UpdateVersion(version Version) error {
	_, err := exec.Command("git", "tag", "-a", version.String(), "-m", "''").CombinedOutput()
	if err != nil {
		return err
	}

	current, err := this.CurrentVersion()
	if err != nil {
		return err
	}

	return assert.So(current, should.Resemble, version).Error()
}
