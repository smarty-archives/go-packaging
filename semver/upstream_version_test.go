package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestUpstreamVersionFixture(t *testing.T) {
	gunit.Run(new(UpstreamVersionFixture), t)
}

type UpstreamVersionFixture struct {
	*gunit.Fixture
}

func (this *UpstreamVersionFixture) assertParseFailure(input string) {
	version, err := ParseUpstream(input)
	this.So(version, should.Resemble, UpstreamVersion{})
	this.So(err, should.NotBeNil)
}

func (this *UpstreamVersionFixture) assertParseSuccess(input, upstream string, revision int) {
	version, err := ParseUpstream(input)
	this.So(version.String(), should.Equal, UpstreamVersion{Version: upstream, Revision: revision}.String())
	this.So(err, should.BeNil)
}

func (this *UpstreamVersionFixture) TestParsing() {
	this.assertParseFailure("")
	this.assertParseFailure("1.2.3.4")
	this.assertParseFailure("1")
	this.assertParseFailure("helloworld")
	this.assertParseFailure("1.b.3")
	this.assertParseFailure("1.2.3-a")
	this.assertParseSuccess("1.2.3-4", "1.2.3", 4)
	this.assertParseSuccess(" \n\t1.2.3-4\t\n ", "1.2.3", 4)
	this.assertParseSuccess("1.2.3-4-1-ab5def", "1.2.3", 4)
	this.assertParseSuccess("1.2.3-rc4-5-1-ab5def", "1.2.3-rc4", 5)
}

func (this *UpstreamVersionFixture) TestDisplay() {
	version := UpstreamVersion{Version: "hello, world!", Revision: 42}
	this.So(version.String(), should.Equal, "hello, world!-42")
}

func (this *UpstreamVersionFixture) TestPatchRemainsUnchangedIfNotDirty() {
	version := UpstreamVersion{Version: "1.2.3", Revision: 4}
	incremented := version.Increment()
	this.So(version, should.Resemble, incremented)
	this.So(incremented.String(), should.Equal, "1.2.3-4")
}

func (this *UpstreamVersionFixture) TestPatchIncrementsWhenDirty() {
	version := UpstreamVersion{Version: "1.2.3", Revision: 4, dirty: true}
	incremented := version.Increment()
	this.So(version, should.NotResemble, incremented)
	this.So(incremented.String(), should.Equal, "1.2.3-5")
}

func (this *UpstreamVersionFixture) TestPatchOnlyIncrementsOnce() {
	version := UpstreamVersion{Version: "1.2.3", Revision: 4, dirty: true}
	incremented := version.Increment().Increment().Increment()
	this.So(version, should.NotResemble, incremented)
	this.So(incremented.String(), should.Equal, "1.2.3-5")
}
