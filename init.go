package main

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func InitDotBible() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Join(home, ".bible")

	_, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Initializing .bible...")
			err = os.Mkdir(baseDir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			versionsDir := filepath.Join(baseDir, "versions")
			err = os.Mkdir(versionsDir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			for _, v := range versions {
				wg.Add(1)
				go InitVersion(v)
			}
			wg.Wait()
			fmt.Println(".bible initialization completed.")
		}
	}

}

func InitVersion(version Version) {
	defer wg.Done()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Join(home, ".bible")
	versionsDir := filepath.Join(baseDir, "versions")

	// Create baseDir/versions/{version.name}
	versionDir := filepath.Join(versionsDir, version.Name)
	os.Mkdir(versionDir, 0755)

	// Download the version
	resp, err := http.Get(version.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	split := strings.Split(version.Url, ".")
	name := filepath.Join(versionDir, version.Name+"."+split[len(split)-1])
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize progress bar
	bar := pb.Full.Start(size)

	// Create a proxy reader to track progress
	reader := bar.NewProxyReader(resp.Body)

	// Write the file
	_, err = io.Copy(file, reader)
	if err != nil {
		log.Fatal(err)
	}

	// Finish the progress bar
	bar.Finish()

	log.Printf("Version %s initialization completed.\n", version.Name)
}
