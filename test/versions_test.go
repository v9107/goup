package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/v3nkat3shk/goup/src"
)

func TestLocalVersion(t *testing.T) {
	version, err := src.GetLocalVersion()

	if err != nil {
		t.Error(err.Error())
	}

	goVersion, err := exec.Command("go", "version").Output()

	if err != nil {
		t.Error(err.Error())
	}

	localVersion := strings.Split(strings.TrimSpace(strings.Replace(string(goVersion), "go version", "", 1)), " ")

	if version.Version == localVersion[0] {
		t.Errorf("fetch version \"%s\" != \"%s\" actual version", version.Version, localVersion[0])
	}
}
