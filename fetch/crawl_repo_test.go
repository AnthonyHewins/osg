package fetch

import (
	"path/filepath"
	"testing"
	"runtime"
	"reflect"
)

var correct_file_list = []string{
	// these files exist but should be whitelisted, which is why they're commented out,
	// but wanted to show them explicitly for documentation:
	//
	// "Cargo.toml",                    <-----  None of these should appear.
	// ".gitignore",                    <-----
	// "Cargo.lock",                    <-----
	// "util/Cargo.toml",               <-----  (good example of parent dir is irrelevant if it's whitelisted via filename)

	"src/main.rs",
	"src/keygen.rs",
	"src/ftp.rs",
	"src/socket/mod.rs",
	"src/socket/io.rs",
	"src/socket/internals.rs",
	"src/listener/mod.rs",
	"util/src/lib.rs",
	"util/src/fs.rs",
	"util/src/net.rs",
	"tmp.txt",
}

// Make sure you have the fixtures for the tests so you can actually run them
const test_fixture = "../fixtures"
const repo_name    = "one-time-pad-socket-master"

var abs_path_fixture string

type Compare = func(chan File, *testing.T) ([]string, []string)

func TestCrawlDir(t *testing.T) {
	abs_path_to_test_dir := get_abs_fixture_path(t)

	var expected []string
	for _, correct_file := range correct_file_list {
		expected = append(
			expected,
			filepath.Join(
				abs_path_to_test_dir,
				correct_file,
			),
		)
	}

	runner(
		&abs_path_to_test_dir,
		t,
		crawl_dir,
		expected,
	)
}

func TestCrawlZip(t *testing.T) {
	abs_path_to_test_dir := get_abs_fixture_path(t) + ".zip"

	runner(
		&abs_path_to_test_dir,
		t,
		crawl_zip,
		build_expected_compressed_file(),
	)
}

func TestCrawlTar(t *testing.T) {
	abs_path_to_test_dir := get_abs_fixture_path(t) + ".tar.gz"

	runner(
		&abs_path_to_test_dir,
		t,
		crawl_tar,
		build_expected_compressed_file(),
	)
}

func get_func_name(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func runner(path *string, t *testing.T, lambda Crawlable, expected []string) {
	files_pipe := make(chan File, 30)
	errors     := make(chan error, 6)

	lambda(path, files_pipe, errors)

	close(files_pipe)
	close(errors)

	var not_present_in_expected []string

	for f := range files_pipe {
		found_it := -1
		for i, filename := range expected {
			if f.Name == filename {
				found_it = i
				break
			}
		}

		if found_it != -1 {
			// this removes element i; since we've now seen it, we can check it off our list
			n := len(expected)
			expected[found_it] = expected[n-1]
			expected = expected[:n-1]
		} else {
			not_present_in_expected = append(not_present_in_expected, f.Name)
		}
	}

	for _, filename := range expected {
		t.Errorf("%v was not scanned running %v and should have been", filename, get_func_name(lambda))
	}

	for _, filename := range not_present_in_expected {
		t.Errorf("%v was either scanned too many times or was an arbitrary unexpected file in %v run", filename, get_func_name(lambda))
	}
}

func build_expected_compressed_file() []string {
	var checklist []string

	for _, correct_file := range correct_file_list {
		checklist = append(
			checklist,
			filepath.Join(repo_name, correct_file),
		)
	}

	return checklist
}

func get_abs_fixture_path(t *testing.T) string {
	if abs_path_fixture == "" {
		abs_path, err := filepath.Abs(test_fixture)
		if err != nil {
			t.Fatalf("couldn't run tests because couldn't resolve full path of %v: %v", test_fixture, err)
		}

		abs_path_fixture = filepath.Join(abs_path, repo_name)
	}

	return abs_path_fixture
}
