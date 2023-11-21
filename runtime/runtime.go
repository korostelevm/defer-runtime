package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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

func get_process_status(procfolder string) map[string]string {
	contents := read_file(procfolder + "/status")
	status := make(map[string]string)
	for _, line := range contents {
		// split the line into key and value
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			status[key] = value
		}
	}

	return status

}

func main() {
	// list all files in current directory
	// enc := json.NewEncoder(os.Stdout)
	for true {

		files, err := os.ReadDir("/proc")
		if err != nil {
			log.Fatal(err)
		}
		// processes := make(map[string]map[string]string)
		processes := []string{}
		// for true {
		for _, file := range files {
			process, _ := regexp.MatchString(`^\d+$`, file.Name())
			if process {
				processes = append(processes, file.Name())
			}
		}

		procs := make(map[string]map[string]string)
		for _, process := range processes {
			path := "/proc/" + process
			status := get_process_status(path)
			// fmt.Println(file.Name(), status["Name"])
			procs[status["Pid"]] = map[string]string{
				"Name": status["Name"],
				"Pid":  status["Pid"],
			}
		}

		j, _ := json.MarshalIndent(procs, "", "  ")
		fmt.Println("\033[2J")
		println(string(j))
	}

}
