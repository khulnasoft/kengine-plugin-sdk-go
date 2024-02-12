package useragent

import (
	"errors"
	"regexp"
)

var (
	userAgentRegex   = regexp.MustCompile(`^Khulnasoft/([0-9]+\.[0-9]+\.[0-9]+(?:-[a-zA-Z0-9]+)?) \(([a-zA-Z0-9]+); ([a-zA-Z0-9]+)\)$`)
	errInvalidFormat = errors.New("invalid user agent format")
)

// UserAgent represents a Khulnasoft user agent.
// Its format is "Khulnasoft/<version> (<os>; <arch>)"
// Example: "Khulnasoft/7.0.0-beta1 (darwin; amd64)", "Khulnasoft/10.0.0 (windows; x86)"
type UserAgent struct {
	kengineVersion string
	arch           string
	os             string
}

// New creates a new UserAgent.
// The version must be a valid semver string, and the os and arch must be valid strings.
func New(kengineVersion, os, arch string) (*UserAgent, error) {
	ua := &UserAgent{
		kengineVersion: kengineVersion,
		os:             os,
		arch:           arch,
	}

	return Parse(ua.String())
}

// Parse creates a new UserAgent from a string.
func Parse(s string) (*UserAgent, error) {
	matches := userAgentRegex.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, errInvalidFormat
	}

	return &UserAgent{
		kengineVersion: matches[1],
		os:             matches[2],
		arch:           matches[3],
	}, nil
}

// Empty creates a new UserAgent with default values.
func Empty() *UserAgent {
	return &UserAgent{
		kengineVersion: "0.0.0",
		os:             "unknown",
		arch:           "unknown",
	}
}

func (ua *UserAgent) KengineVersion() string {
	return ua.kengineVersion
}

func (ua *UserAgent) String() string {
	return "Khulnasoft/" + ua.kengineVersion + " (" + ua.os + "; " + ua.arch + ")"
}
