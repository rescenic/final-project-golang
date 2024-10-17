package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "./"

	// Output text file where all content will be saved
	outputFile := "all_files_output.txt"
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer f.Close()

	// Walk through the root directory
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Process both .go and .sql files
		if strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), ".sql") {
			// Write the file path as a header
			header := fmt.Sprintf("// %s\n", path)
			_, err := f.WriteString(header)
			if err != nil {
				return err
			}

			// Read the file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Write the file content
			_, err = f.Write(content)
			if err != nil {
				return err
			}

			// Write a separator line
			_, err = f.WriteString("\n\n")
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through the project:", err)
	} else {
		fmt.Println("All files merged into:", outputFile)
	}
}
