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

func (this *UpstreamVersionFixture) assertParseFailure(base, raw string) {
	parsed, err := ParseUpstream(base, raw)
	this.So(parsed, should.Resemble, UpstreamVersion{})
	this.So(err, should.NotBeNil)
}

func (this *UpstreamVersionFixture) assertParseSuccess(base, raw, expected string) {
	parsed, err := ParseUpstream(base, raw)
	this.So(parsed.String(), should.Equal, expected)
	this.So(err, should.BeNil)
}

func (this *UpstreamVersionFixture) TestParsing() {
	this.assertParseFailure("", "")
	this.assertParseFailure("", "1-1")
	this.assertParseSuccess("1.2.3", "", "1.2.3-0")
	this.assertParseSuccess("1.2.3", "1.2.3-1", "1.2.3-1")
	this.assertParseSuccess("1.2.3", "1.2.33", "1.2.3-0")
	this.assertParseSuccess("1.2.3", "1.2.33-1", "1.2.3-0")
	this.assertParseSuccess("4.5.6", "1.2.3-1", "4.5.6-0")
	this.assertParseSuccess("7.8.9", "7.8.9-0-17-abcdef", "7.8.9-1")
	this.assertParseSuccess("base", "base-1", "base-1")
	this.assertParseSuccess("base-rc1", "base-rc1-1", "base-rc1-1")
	this.assertParseSuccess("base-rc1", "base-rc1-12-abcdef", "base-rc1-13")
	this.assertParseSuccess("base-rc2", "base-rc1-12-abcdef", "base-rc2-0")
	this.assertParseFailure("1.2.3", "1.2.3-a")
	this.assertParseFailure("1.2.3", "1.2.3-a-abcdef")
	this.assertParseSuccess("  1.2.3  ", "  1.2.3-1  ", "1.2.3-1")
}

func (this *UpstreamVersionFixture) TestDisplay() {
	this.So(UpstreamVersion{Base: "1.2.3", Revision: 42}.String(), should.Equal, "1.2.3-42")
	this.So(UpstreamVersion{Base: "1.2.3-rc1", Revision: 42}.String(), should.Equal, "1.2.3-rc1-42")
	this.So(UpstreamVersion{Base: "base", Revision: 2}.String(), should.Equal, "base-2")
}
