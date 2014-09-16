package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersionInfo(t *testing.T) {
	Convey("When parsing:", t, func() {
		Convey("When the provided version is empty, it should return an error", func() {
			version, err := ParseVersion("")
			So(version, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("When the provided version has too many parts", func() {
			version, err := ParseVersion("1.2.3.4")
			So(version, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("When the provided version has too few parts", func() {
			version, err := ParseVersion("1")
			So(version, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("When the provided version contains a simple, non-numeric word", func() {
			version, err := ParseVersion("helloworld")
			So(version, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("When the provided version non-numeric parts", func() {
			version, err := ParseVersion("1.b.3")
			So(version, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("When the provided version is a valid, 3-part version", func() {
			version, err := ParseVersion("1.2.3")
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 3)
			So(err, ShouldBeNil)
		})
		Convey("When the provided version is a valid, 2-part version", func() {
			version, err := ParseVersion("1.2")
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 0)
			So(err, ShouldBeNil)
		})
		Convey("When the provided version is a valid but contains whitespace", func() {
			version, err := ParseVersion(" 1 . 2 . 3 ")
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 3)
			So(err, ShouldBeNil)
		})
	})
	Convey("When representing the version as a string", t, func() {
		Convey("It should display all the parts", func() {
			version123 := VersionInfo{Major: 1, Minor: 2, Patch: 3}
			So(version123.String(), ShouldEqual, "1.2.3")

			version456 := VersionInfo{Major: 4, Minor: 5, Patch: 6}
			So(version456.String(), ShouldEqual, "4.5.6")
		})
	})
}
