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
			So(version, ShouldNotBeNil)
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 3)
			So(err, ShouldBeNil)
		})
		Convey("When the provided version is a valid, 2-part version", func() {
			version, err := ParseVersion("1.2")
			So(version, ShouldNotBeNil)
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 0)
			So(err, ShouldBeNil)
		})
		Convey("When the provided version is a valid but contains whitespace", func() {
			version, err := ParseVersion(" 1 . 2 . 3 ")
			So(version, ShouldNotBeNil)
			So(version.Major, ShouldEqual, 1)
			So(version.Minor, ShouldEqual, 2)
			So(version.Patch, ShouldEqual, 3)
			So(err, ShouldBeNil)
		})
		Convey("When the provided version is a git tag that has a commit distance", func() {
			version, err := ParseVersion("1.2.3-1-ab5def")
			So(version, ShouldNotBeNil)
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

	Convey("When incrementing the patch number", t, func() {
		Convey(`If the version is NOT marked as "dirty", it should not increment any number.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: false}
			incremented := version.IncrementPatch()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.2.3")
		})
		Convey(`If the version is marked as "dirty", it should increment the patch number.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementPatch()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.2.4")
		})
		Convey(`If the version is incremented multiple times, it should only increment once.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementPatch().IncrementPatch().IncrementPatch()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.2.4")
		})
	})

	Convey("When incrementing the minor number", t, func() {
		Convey(`If the version is NOT marked as "dirty", it should not increment any number.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: false}
			incremented := version.IncrementMinor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.2.3")
		})
		Convey(`If the version is marked as "dirty", it should increment the minor number and reset the patch number.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementMinor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.3.0")
		})
		Convey(`If the version is incremented multiple times, it should only increment once.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementMinor().IncrementMinor().IncrementMinor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.3.0")
		})
	})

	Convey("When incrementing the major number", t, func() {
		Convey(`If the version is NOT marked as "dirty", it should not increment any number.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: false}
			incremented := version.IncrementMajor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldPointTo, incremented)
			So(incremented.String(), ShouldEqual, "1.2.3")
		})
		Convey(`If the version is marked as "dirty", it should increment the major number and reset the other numbers.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementMajor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "2.0.0")
		})
		Convey(`If the version is incremented multiple times, it should only increment once.`, func() {
			version := &VersionInfo{Major: 1, Minor: 2, Patch: 3, dirty: true}
			incremented := version.IncrementMajor().IncrementMajor().IncrementMajor()
			So(incremented, ShouldNotBeNil)
			So(version, ShouldNotPointTo, incremented)
			So(incremented.String(), ShouldEqual, "2.0.0")
		})
	})
}
