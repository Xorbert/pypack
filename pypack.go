package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func createPackage(name string, prompt bool) {
	path := filepath.Join(name, "__init__.py")
	if _, err := os.Stat(path); err == nil {
		if prompt {
			fmt.Printf("%s/__init__.py already exists. Overwrite it? (y/n) ", name)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			response := scanner.Text()
			if response != "y" {
				fmt.Println("Aborting.")
				os.Exit(1)
			}
		} else {
			return
		}
	}
	err := os.MkdirAll(name, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", name, err)
		os.Exit(1)
	}
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Printf("Successfully created package %s.\n", name)
}

func main() {
	var prompt = false
	var name string
	if len(os.Args) > 2 && os.Args[1] == "-i" {
		prompt = true
		name = os.Args[2]
	} else if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		fmt.Println("Usage: script.exe [-i] packagename")
		os.Exit(1)
	}
	createPackage(name, prompt)
}
