package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestVersionInfoFixture(t *testing.T) {
	gunit.Run(new(VersionInfoFixture), t)
}

type VersionInfoFixture struct {
	*gunit.Fixture
}

func (this *VersionInfoFixture) assertParseFailure(input string) {
	version, err := ParseVersion(input)
	this.So(version, should.BeNil)
	this.So(err, should.NotBeNil)
}
func (this *VersionInfoFixture) assertParseSuccess(input string, major, minor, patch int) {
	version, err := ParseVersion(input)
	this.So(err, should.BeNil)
	if this.So(version, should.NotBeNil) {
		this.So(version.Major, should.Equal, major)
		this.So(version.Minor, should.Equal, minor)
		this.So(version.Patch, should.Equal, patch)
	}
}

func (this *VersionInfoFixture) TestParsing() {
	this.assertParseFailure("")
	this.assertParseFailure("1.2.3.4")
	this.assertParseFailure("1")
	this.assertParseFailure("helloworld")
	this.assertParseFailure("1.b.3")
	this.assertParseSuccess("1.2.3", 1, 2, 3)
	this.assertParseSuccess("1.2", 1, 2, 0)
	this.assertParseSuccess("1.2", 1, 2, 0)
	this.assertParseSuccess(" 1 . 2 . 3 ", 1, 2, 3)
	this.assertParseSuccess("1.2.3-1-ab5def", 1, 2, 3)
}

func (this *VersionInfoFixture) TestDisplay() {
	version123 := VersionInfo{Major: 1, Minor: 2, Patch: 3}
	this.So(version123.String(), should.Equal, "1.2.3")

	version456 := VersionInfo{Major: 4, Minor: 5, Patch: 6}
	this.So(version456.String(), should.Equal, "4.5.6")
}

func (this *VersionInfoFixture) TestPatchRemainsUnchangedIfNotDirty() {
	version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: false}
	incremented := version.IncrementPatch()
	this.So(incremented, should.NotBeNil)
	this.So(version, should.PointTo, incremented)
	this.So(incremented.String(), should.Equal, "1.2.3")
}

func (this *VersionInfoFixture) TestPatchIncrementsWhenDirty() {
	version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
	incremented := version.IncrementPatch()
	this.So(incremented, should.NotBeNil)
	this.So(version, should.NotPointTo, incremented)
	this.So(incremented.String(), should.Equal, "1.2.4")
}

func (this *VersionInfoFixture) TestPatchOnlyIncrementsOnce() {
	version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
	incremented := version.IncrementPatch().IncrementPatch().IncrementPatch()
	this.So(incremented, should.NotBeNil)
	this.So(version, should.NotPointTo, incremented)
	this.So(incremented.String(), should.Equal, "1.2.4")
}
