package fetch

import (
	"testing"
	"io/ioutil"
	"os"
)

type testcase struct {
	name string
	is_whitelisted bool
}

var test_strings = []testcase{
	{ "analytics_engine.tar", false },
	{ "big chungus",          false },
	{ "asd",                  false },
	{ "media*.mp3",            true },
}

func TestWhitelisted(t *testing.T) {
	run_against := func(name *string, should_be_whitelisted bool) {
		f, err := ioutil.TempFile("", *name)
		if err != nil {
			t.Errorf("unable to create file %v for testing: %v", name, err)
			return
		}

		defer f.Close()
		defer os.Remove(f.Name())

		metadata, err := f.Stat()
		if err != nil {
			t.Errorf("unable to test %v, couldn't get metadata", f.Name())
			return
		}

		is_whitelisted := whitelisted(metadata)
		if is_whitelisted != should_be_whitelisted { // XOR?
			t.Errorf("whitelisted status should be %v, but wasn't for %v", should_be_whitelisted, *name)
		}
	}

	for _, t := range test_strings { run_against(&t.name, t.is_whitelisted) }
}
