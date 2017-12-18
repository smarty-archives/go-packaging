package main

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestVersionIncrementationFixture(t *testing.T) {
	assert := assertions.New(t)
	assert.So(version(1, 2, 3, true).IncrementMajor(), should.Resemble, version(2, 0, 0, false))
	assert.So(version(1, 2, 3, true).IncrementMinor(), should.Resemble, version(1, 3, 0, false))
	assert.So(version(1, 2, 3, true).IncrementPatch(), should.Resemble, version(1, 2, 4, false))

	assert.So(version(1, 2, 3, true).Increment("mAjOr"), should.Resemble, version(2, 0, 0, false))
	assert.So(version(1, 2, 3, true).Increment("MiNoR"), should.Resemble, version(1, 3, 0, false))
	assert.So(version(1, 2, 3, true).Increment("PATCH"), should.Resemble, version(1, 2, 4, false))
	assert.So(version(1, 2, 3, true).Increment(""), should.Resemble, version(1, 2, 4, false))
}

func TestVersionString(t *testing.T) {
	assertions.New(t).So(version(1, 2, 3, false).String(), should.Equal, "1.2.3")
	assertions.New(t).So(version(1, 2, 3, true).String(), should.Equal, "1.2.3")
}

func version(major, minor, patch int, dirty bool) Version {
	return Version{Major: major, Minor: minor, Patch: patch, Dirty: dirty}
}
