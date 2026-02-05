package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

var (
	binName  = "test"
	cmdPath  string
	exitCode int
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build %s: %s", binName, err)
		os.Exit(1)
	}

	var err error
	cmdPath, err = filepath.Abs(binName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get absolute path to %s: %s", binName, err)
		os.Exit(1)
	}

	exitCode = m.Run()

	os.Remove(binName)
	os.Exit(exitCode)
}

func TestNoArgs(t *testing.T) {
	cmd := exec.Command(cmdPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	expected := "-"
	if !strings.Contains(out.String(), expected) {
		t.Errorf("expected to contain %q, got %q", expected, out.String())
	}
}

func TestCorrectFlags(t *testing.T) {
	testCases := []struct {
		args     []string
		expected string
	}{
		{[]string{"-v"}, "version"},
		{[]string{"-h"}, "Usage:"},
		{[]string{"-u"}, "_"},
		{[]string{"-c"}, "ch)"},
		{[]string{"-w"}, "Usage:"},
		{[]string{"-w", "4"}, "-"},
		{[]string{"-w", "0"}, "must be"},
		{[]string{"-w", "-1"}, "must be"},
		{[]string{"-w", "a"}, "Usage:"},
		{[]string{"-w", "!"}, "Usage:"},
		{[]string{"-d"}, "Usage:"},
		{[]string{"-d", "4"}, "-"},
		{[]string{"-d", "0"}, "must be"},
		{[]string{"-d", "-1"}, "must be"},
		{[]string{"-d", "a"}, "Usage:"},
		{[]string{"-d", "!"}, "Usage:"},
	}

	for _, tc := range testCases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			cmd := exec.Command(cmdPath, tc.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()

			if !strings.Contains(out.String(), tc.expected) {
				t.Errorf("expected to contain %q, got %q", tc.expected, out.String())
			}
		})
	}
}

func TestWords(t *testing.T) {
	testCases := []struct {
		args     []string
		expected int
	}{
		{[]string{"-w", "1"}, 1},
		{[]string{"-w", "4"}, 4},
		{[]string{"-w", "7"}, 7},
	}

	for _, tc := range testCases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			cmd := exec.Command(cmdPath, tc.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()

			count := strings.Count(out.String(), "-")

 			if count != tc.expected {
				t.Errorf("expected %d dashes, got %d", tc.expected, count)
			}
		})
	}
}

func TestDigits(t *testing.T) {
	testCases := []struct {
		args     []string
		expected int
	}{
		{[]string{"-d", "1"}, 1},
		{[]string{"-d", "4"}, 4},
		{[]string{"-d", "7"}, 7},
	}

	for _, tc := range testCases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			cmd := exec.Command(cmdPath, tc.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Run()

			count := 0
			for _, r := range out.String() {
				if unicode.IsDigit(r) {
					count++
				}
			}

 			if count != tc.expected {
				t.Errorf("expected %d digits, got %d", tc.expected, count)
			}
		})
	}
}

func TestWrongFlag(t *testing.T) {
	cmd := exec.Command(cmdPath, "-wrong")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	expected := "Usage:"
	if !strings.Contains(out.String(), expected) {
		t.Errorf("expected to contain %q, got %q", expected, out.String())
	}
}

// unit tests

func TestGetDigits(t *testing.T) {
	digits := getNum(4)
	if len(digits) != 4 {
		t.Errorf("expected 4 digits, got %d", len(digits))
	}
	for _, r := range digits {
		if !unicode.IsDigit(r) {
			t.Errorf("expected digit, got %q", r)
		}
	}
}

func TestGetDashPhrase(t *testing.T) {
	phrase := getDashPhrase(3, 2)
	parts := strings.Split(phrase, "-")
	if len(parts) != 4 {
		t.Errorf("expected 4 parts separated by dash, got %d", len(parts))
	}
	// make sure first 3 parts are words
	for i := 0; i < 3; i++ {
		if len(parts[i]) == 0 {
			t.Errorf("expected non-empty word at part %d", i)
		}
	}
	// make sure last part is 2 digits
	if len(parts[3]) != 2 {
		t.Errorf("expected 2 digits, got %d", len(parts[3]))
	}
}

func TestGetNamedDashPhrase(t *testing.T) {
	phrase := getNamedDashPhrase(2, 3)
	parts := strings.Split(phrase, "-")
	if len(parts) != 3 {
		t.Errorf("expected 3 parts separated by dash, got %d", len(parts))
	}
	// make sure first 2 parts are words
	for i := 0; i < 2; i++ {
		if len(parts[i]) == 0 {
			t.Errorf("expected non-empty word at part %d", i)
		}
	}
	// make sure last part is 3 digits
	if len(parts[2]) != 3 {
		t.Errorf("expected 3 digits, got %d", len(parts[2]))
	}
}

func TestGetPhrase(t *testing.T) {
	phrase := getPhrase(5)
	words := strings.Split(phrase, "-")
	if len(words) != 5 {
		t.Errorf("expected 5 words, got %d", len(words))
	}
}

func TestGetNamedPhrase(t *testing.T) {
	phrase := getNamedPhrase(4)
	words := strings.Split(phrase, "-")
	if len(words) != 4 {
		t.Errorf("expected 4 words, got %d", len(words))
	}
}

func TestGetWord(t *testing.T) {
	word := getWord()
	if len(word) == 0 {
		t.Errorf("expected non-empty word")
	}
}

func TestGetName(t *testing.T) {
	name := getName()
	if len(name) == 0 {
		t.Errorf("expected non-empty name")
	}
}

func TestVersion(t *testing.T) {
	output := captureOutput(Version)
	if !strings.Contains(output, "version") {
		t.Errorf("expected output to contain 'version', got %q", output)
	}
}

func TestUsage(t *testing.T) {
	output := captureOutput(Usage)
	if !strings.Contains(output, "Usage:") {
		t.Errorf("expected output to contain 'Usage:', got %q", output)
	}
}

// captureOutput captures the output of a function that prints to stdout.
func captureOutput(f func()) string {

	old := os.Stdout
	old_err := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w


	f()

	w.Close()
	os.Stdout = old
	os.Stderr = old_err

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func captureOutputWithError(f func() error) (string, error) {

	old := os.Stdout
	old_err := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w


	err := f()

	w.Close()
	os.Stdout = old
	os.Stderr = old_err

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String(), err
}


func TestRun(t *testing.T) {

		// default flags
		nm3 := 3
		nm0 := 0
		blf := false
		blt := true

		digitsFlag = &nm3
		wordsFlag = &nm3
		versionFlag = &blf
		helpFlag = &blf
		countFlag = &blf
		nameFlag = &blf
		underscoreFlag = &blf

	t.Run("NoArgs", func(t *testing.T) {


		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "-") {
			t.Errorf("expected output to contain '-', got %q", output)
		}
	})

	t.Run("InvalidWords", func(t *testing.T) {

		wordsFlag = &nm0

		output,err := captureOutputWithError(run)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "number of words must be at least 1") {
			t.Errorf("expected error to contain 'number of words must be at least 1', got %q", err.Error())
		}
		if !strings.Contains(output, "") {
			t.Errorf("expected output to be empty, got %q", output)
		}
	})

	t.Run("InvalidDigits", func(t *testing.T) {

		wordsFlag = &nm3
		digitsFlag = &nm0

		output,err := captureOutputWithError(run)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "number of digits must be at least 1") {
			t.Errorf("expected error to contain 'number of digits must be at least 1', got %q", err.Error())
		}
		if !strings.Contains(output, "") {
			t.Errorf("expected output to be empty, got %q", output)
		}
	})

	t.Run("VersionFlag", func(t *testing.T) {

		versionFlag = &blt
		helpFlag = &blf
		digitsFlag = &nm3
		wordsFlag = &nm3

		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "version") {
			t.Errorf("expected output to contain 'version', got %q", output)
		}
	})

	t.Run("CountFlag", func(t *testing.T) {

		versionFlag = &blf
		helpFlag = &blf
		digitsFlag = &nm3
		wordsFlag = &nm3
		countFlag = &blt

		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "ch)") {
			t.Errorf("expected output to contain 'ch)', got %q", output)
		}
	})

	t.Run("UnderscoreFlag", func(t *testing.T) {

		versionFlag = &blf
		helpFlag = &blf
		digitsFlag = &nm3
		wordsFlag = &nm3
		underscoreFlag = &blt

		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "_") {
			t.Errorf("expected output to contain '_', got %q", output)
		}
	})


	t.Run("NameFlag", func(t *testing.T) {
		
		versionFlag = &blf
		helpFlag = &blf
		digitsFlag = &nm3
		wordsFlag = &nm3
		nameFlag = &blt
		underscoreFlag = &blf

		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "-") {
			t.Errorf("expected output to contain '-', got %q", output)
		}
	})

	t.Run("HelpFlag", func(t *testing.T) {

		versionFlag = &blf
		helpFlag = &blt
		digitsFlag = &nm3
		wordsFlag = &nm3


		output,err := captureOutputWithError(run)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !strings.Contains(output, "Usage:") {
			t.Errorf("expected output to contain 'Usage:', got %q", output)
		}
	})


}
