package main

import (
	"fmt"
	"log"

	"github.com/v9107/goup/src"
)

var update = "no"

func main() {
	versions, err := src.GetVersions(src.LATEST_VERSION_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	canBeUpdated, err := versions.CheckForUpdates()

	if err != nil {
		log.Fatal(err.Error())
	}

	if !canBeUpdated {
		log.Println("Version is up to date")
		return
	}

	log.Default().Printf("New version is available %s -> %s\n", versions.LocalVersion.Version, versions.LatestVersion.Version)
	fmt.Print("Would you like to update [yes/no] (default is no): ")
	fmt.Scan(&update)

	if update != "yes" {
		return
	}

	if err := versions.DownloadLatestVersion(src.DOWNLOAD_URL); err != nil {
		log.Fatal(err.Error())
	}
}
