package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func InitDotBible() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Join(home, ".bible")

	_, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
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
				// Create baseDir/versions/{version.name}
				versionDir := filepath.Join(versionsDir, v.Name)
				os.Mkdir(versionDir, 0755)

				// Download the version
				resp, err := http.Get(v.Url)
				if err != nil {
					log.Fatal(err)
				}

				split := strings.Split(v.Url, ".")
				name := filepath.Join(versionDir, v.Name+"."+split[len(split)-1])
				file, err := os.Create(name)
				if err != nil {
					log.Fatal(err)
				}

				// Write the file
				_, err = io.Copy(file, resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				if err := file.Close(); err != nil {
					log.Fatal(err)
				}

				if err := resp.Body.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
