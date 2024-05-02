package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	filenames := make(chan string)

	dir := os.Args[1]

	go getFiles(dir, filenames)
	buffer := make([]string, 0)
	for filename := range filenames {
		buffer = append(buffer, filename)
	}
	fmt.Println(buffer)

	end := time.Now()

	fmt.Printf("Elapsed Time: %v\n", end.Sub(start))

	fmt.Printf("Total CPU: %d\n", runtime.NumCPU())

	fmt.Printf("Total Goroutines: %d\n", runtime.NumGoroutine())
}

func getFiles(dir string, fileNames chan<- string) {
	defer close(fileNames)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileNames <- path
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
