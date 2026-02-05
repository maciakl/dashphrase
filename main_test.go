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

	expected := "Usage:"
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
		{[]string{"--version"}, "version"},
		{[]string{"-h"}, "Usage:"},
		{[]string{"--help"}, "Usage:"},
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