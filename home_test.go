package home

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/gotoxu/assert"
)

func BenchmarkDir(b *testing.B) {
	// warmups
	for i := 0; i < 10; i++ {
		Dir()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dir()
	}
}

func TestDir(t *testing.T) {
	u, err := user.Current()
	assert.Nil(t, err)

	dir, err := Dir()
	assert.Nil(t, err)
	assert.DeepEqual(t, dir, u.HomeDir)
}

func TestExpand(t *testing.T) {
	u, err := user.Current()
	assert.Nil(t, err)

	cases := []struct {
		Input  string
		Output string
		Err    bool
	}{
		{
			"/foo",
			"/foo",
			false,
		},

		{
			"~/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},

		{
			"",
			"",
			false,
		},

		{
			"~",
			u.HomeDir,
			false,
		},

		{
			"~foo/foo",
			"",
			true,
		},
	}

	for _, tc := range cases {
		actual, err := Expand(tc.Input)
		if !tc.Err {
			assert.Nil(t, err)
		}

		assert.DeepEqual(t, actual, tc.Output)
	}

	DisableCache = true
	defer func() { DisableCache = false }()
	defer patchEnv("HOME", "/custom/path/")()
	expected := filepath.Join("/", "custom", "path", "foo/bar")
	actual, err := Expand("~/foo/bar")

	assert.Nil(t, err)
	assert.DeepEqual(t, actual, expected)
}

func patchEnv(key, value string) func() {
	bck := os.Getenv(key)
	deferFunc := func() {
		os.Setenv(key, bck)
	}

	os.Setenv(key, value)
	return deferFunc
}
