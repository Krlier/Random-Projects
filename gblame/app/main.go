package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Output holds all the JSON output
type Output struct {
	Issues []Issue `json:"Issues"`
}

// Issue is a struct to hold all the details about a security issue found
type Issue struct {
	Details  string `json:"details"`
	FilePath string `json:"file"`
	Line     string `json:"line"`
}

func main() {
	// Reads the resulting JSON of the static analysis tool
	file, err := ioutil.ReadFile("results.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshalls the json into the data structure
	output := Output{}
	json.Unmarshal([]byte(file), &output)

	// Creates the git blame script
	f, err := os.Create("gblame.sh")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("#!/bin/bash\n")

	// Adds a new git blame command for each security vulnerability entry
	for _, issue := range output.Issues {
		line := fmt.Sprintf("%s,+1", issue.Line)
		dir := filepath.Dir(issue.FilePath)
		command := fmt.Sprintf("cd %s && git blame -L %s %s | awk -F \" \" '{print $2 $3}'\n", dir, line, issue.FilePath)

		f.WriteString(command)

		exec.Command("chmod", "+x", "gblame.sh").Output()
	}

	// Execute the script and print the results
	cmd, _ := exec.Command("./gblame.sh").Output()
	fmt.Printf("%s", cmd)

}
