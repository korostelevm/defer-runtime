package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func read_file(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Make sure to close the file when you're done

	// Create a new Scanner to read the file
	scanner := bufio.NewScanner(file)

	contents := make([]string, 0)
	// Loop over all lines in the file
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return contents
}

func get_process_stats(procfolder string) []string {
	contents := read_file(procfolder + "/stat")
	return contents

}

func main() {
	// list all files in current directory
	files, err := os.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		process, _ := regexp.MatchString(`^\d+$`, file.Name())
		if process {
			path := "/proc/" + file.Name()
			fmt.Println("process", file.Name())
			status := get_process_status(path)
			fmt.Println(status)

		}

	}

}
