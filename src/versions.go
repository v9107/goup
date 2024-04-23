package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func (v Versions) CheckForUpdates() (bool, error) {
	latestVersion, err := convertVerion(strings.Split(strings.Replace(v.LatestVersion.Version, "go", "", 1), "."))
	if err != nil {
		return false, err
	}

	localVersion, err := convertVerion(strings.Split(strings.Replace(v.LocalVersion.Version, "go", "", 1), "."))
	if err != nil {
		return false, err
	}

	if len(localVersion) != len(latestVersion) {
		return false, fmt.Errorf("cannot be compair local and latest go version")
	}

	for idx, value := range latestVersion {
		if value > localVersion[idx] {
			return true, nil
		}
	}

	return false, nil
}

func (v Versions) DownloadLatestVersion() error {
	version := v.LatestVersion.Version
	os := v.LocalVersion.Os
	arch := v.LocalVersion.Arch

	finalVersion := fmt.Sprintf("%s %s/%s", version, os, arch)
	fmt.Println("Version to Download: ", finalVersion)
	return nil
}

func GetLatestVersion(url string) (APIResponse, error) {
	data := make([]APIResponse, 1)
	res, err := http.Get(url)
	if err != nil {
		return APIResponse{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return APIResponse{}, err

	}

	if err := json.Unmarshal(body, &data); err != nil {
		return APIResponse{}, err
	}

	return data[0], nil
}

func GetLocalVersion() (LocalInstallation, error) {
	goVersion, err := exec.Command("go", "version").Output()

	if err != nil {
		return LocalInstallation{Installed: false}, err
	}

	version := strings.Split(strings.TrimSpace(strings.Replace(string(goVersion), "go version", "", 1)), " ")
	osInfo := strings.Split(version[1], "/")

	if err != nil {
		return LocalInstallation{}, err
	}

	return LocalInstallation{
		Installed: true,
		Version:   version[0],
		Os:        osInfo[0],
		Arch:      osInfo[1],
	}, nil
}
