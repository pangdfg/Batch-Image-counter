package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"batch/image-counter/response"
	s "batch/image-counter/services"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config/config.yml", "config file path")
	flag.Parse()

	cfg, err := s.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	fmt.Println("Input Folder :", cfg.InputFolder)
	fmt.Println("Recursive    :", cfg.Recursive)
	fmt.Println("Workers      :", cfg.Workers)
	fmt.Println()

	files := make(chan string, 100)

	var wg sync.WaitGroup

	counts := map[string]int{
		".jpg":  0,
		".jpeg": 0,
		".png":  0,
		".gif":  0,
		".webp": 0,
		".bmp":  0,
	}

	var folderCount int
	var fileCount int

	var mu sync.Mutex

	counter := func() {
		defer wg.Done()

		for file := range files {
			ext := strings.ToLower(filepath.Ext(file))

			mu.Lock()
			if _, ok := counts[ext]; ok {
				counts[ext]++
			}
			mu.Unlock()
		}
	}

	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go counter()
	}

if cfg.Recursive {

		err = filepath.Walk(cfg.InputFolder,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				if info.IsDir() {
					// Exclude root folder
					if path != cfg.InputFolder {
						folderCount++
					}
					return nil
				}

				fileCount++
				files <- path

				return nil
			})

	} else {

		entries, err := os.ReadDir(cfg.InputFolder)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {

			if entry.IsDir() {
				folderCount++
				continue
			}

			fileCount++
			files <- filepath.Join(cfg.InputFolder, entry.Name())
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	close(files)
	wg.Wait()

	response.PrintConfig(
	cfg.InputFolder,
	cfg.Recursive,
	cfg.Workers,
	)

	totalImages := response.PrintImageCounts(counts)

	response.PrintSummary(
		folderCount,
		fileCount,
		totalImages,
	)
}