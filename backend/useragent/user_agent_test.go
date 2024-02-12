package useragent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type args struct {
		kengineVersion string
		os             string
		arch           string
	}
	tcs := []struct {
		name string
		args args
		want *UserAgent
		err  error
	}{
		{
			name: "valid",
			args: args{
				kengineVersion: "10.2.0",
				os:             "darwin",
				arch:           "amd64",
			},
			want: &UserAgent{
				kengineVersion: "10.2.0",
				os:             "darwin",
				arch:           "amd64",
			},
		},
		{
			name: "invalid (not semver)",
			args: args{
				kengineVersion: "10.2",
				os:             "darwin",
				arch:           "amd64",
			},
			err: errInvalidFormat,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := New(tc.args.kengineVersion, tc.args.os, tc.args.arch)
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.want, got, "New(%v, %v, %v)", tc.args.kengineVersion, tc.args.os, tc.args.arch)
		})
	}
}

func TestParse(t *testing.T) {
	tcs := []struct {
		name      string
		userAgent string
		expected  *UserAgent
		err       error
	}{
		{
			name:      "valid",
			userAgent: "Khulnasoft/10.2.0 (darwin; amd64)",
			expected: &UserAgent{
				kengineVersion: "10.2.0",
				os:             "darwin",
				arch:           "amd64",
			},
		}, {
			name:      "valid (with semver suffix)",
			userAgent: "Khulnasoft/7.0.0-beta1 (darwin; amd64)",
			expected: &UserAgent{
				kengineVersion: "7.0.0-beta1",
				os:             "darwin",
				arch:           "amd64",
			},
		},
		{
			name:      "invalid (missing os + arch)",
			userAgent: "Khulnasoft/7.0.0-beta1",
			err:       errInvalidFormat,
		},
		{
			name:      "invalid (missing arch)",
			userAgent: "Khulnasoft/7.0.0-beta1 (darwin)",
			err:       errInvalidFormat,
		},
		{
			name:      "invalid (missing os)",
			userAgent: "Khulnasoft/7.0.0-beta1 (; amd64)",
			err:       errInvalidFormat,
		},
		{
			name:      "invalid (missing semicolon)",
			userAgent: "Khulnasoft/7.0.0-beta1 (darwin amd64)",
			err:       errInvalidFormat,
		},
		{
			name:      "invalid (not semver)",
			userAgent: "Khulnasoft/10.0 (darwin; amd64)",
			err:       errInvalidFormat,
		},
		{
			name:      "invalid (extra param)",
			userAgent: "Khulnasoft/7.0.0-beta1 (darwin; amd64; linux)",
			err:       errInvalidFormat,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Parse(tc.userAgent)
			require.ErrorIs(t, err, tc.err)
			require.Equalf(t, tc.expected, res, "Parse(%v)", tc.userAgent)
		})
	}
}

func TestEmpty(t *testing.T) {
	expected := &UserAgent{
		kengineVersion: "0.0.0",
		os:             "unknown",
		arch:           "unknown",
	}
	res := Empty()
	require.Equal(t, expected, res)
}
