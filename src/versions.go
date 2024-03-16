package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func (v Versions) CheckForUpdates() (bool, error) {
	latestVersion, err := ConvertVerion(strings.Split(strings.Replace(v.LatestVersion.Version, "go", "", 1), "."))
	if err != nil {
		return false, err
	}

	localVersion, err := ConvertVerion(strings.Split(strings.Replace(v.LocalVersion.Version, "go", "", 1), "."))
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

func (v Versions) DownloadLatestVersion(downloadURL string) error {
	log.Default().Printf("Downloading %s go version", v.LatestVersion.Version)
	urlGenerated := false
	fileName := ""
	for _, file := range v.LatestVersion.Files {
		if file.Os == v.LocalVersion.Os && file.Arch == v.LocalVersion.Arch {
			fileName = file.Filename
			downloadURL = downloadURL + fileName
			urlGenerated = true
		}
	}

	if !urlGenerated {
		return fmt.Errorf("download url not found")
	}

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}

	res, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := io.Copy(out, res.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Downloading completed: %s %dMb\n", fileName, f/int64(math.Pow(10, 6)))
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

	return LocalInstallation{
		Installed: true,
		Version:   version[0],
		Os:        osInfo[0],
		Arch:      osInfo[1],
	}, nil
}
