package main

type Versioning interface {
	CurrentVersion() (Version, error)
	UpdateVersion(Version) error
}
