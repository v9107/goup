package src

import (
	"strconv"
)

const (
	LATEST_VERSION_URL string = "https://go.dev/dl/?mode=json"
	DOWNLOAD_URL       string = "https://storage.googleapis.com/golang/"
)

func GetVersions(url string) (Versions, error) {
	latestVerison, err := GetLatestVersion(url)

	if err != nil {
		return Versions{}, err
	}

	localVersion, err := GetLocalVersion()

	if err != nil {
		return Versions{}, err
	}

	return Versions{
		LatestVersion: latestVerison,
		LocalVersion:  localVersion,
	}, nil
}

func convertVerion(versionArray []string) ([]uint, error) {

	userVersionInt := make([]uint, len(versionArray))

	for idx, value := range versionArray {
		intValue, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return nil, err
		}
		userVersionInt[idx] = uint(intValue)
	}

	return userVersionInt, nil

}
