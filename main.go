package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// setting up main function
func main() {
	// 1 - Accept the dir path from user
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [Dir Path]")
		return
	}
	// Get dit path from the command line argument
	dirPath := os.Args[1]
	fmt.Println("Orginize files in:", dirPath)
	
	// Creating a function a read/get the files from the directory
	files, err := readFiles(dirPath)
	if err != nil {
		fmt.Printf("Error reading Directory: %v\n", err)
		return
	}

	// Loop throughh each file and orgamize them by category

	for _, file := range files {
		if !file.IsDir() {
			category := getCategory(file.Name())
			categoryPath := filepath.Join(dirPath, category)
			err := createDir(dirPath, category)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}

			// Moove the file/S to the category folder
			srcPath := filepath.Join(dirPath, file.Name())
			err = moveFile(srcPath, categoryPath)
			if err != nil {
				fmt.Println("Error moving the file for X reason", err)
			} else {
				fmt.Printf("Moved: %s -> %s\n", file.Name(), category)
			}
		}
	}


}	






// Function to read files from the directory
func readFiles(dirPath string)([]os.DirEntry, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// function to orginize files into different folders based on their extensions
func getCategory(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
		case ".jpg" ,".jpeg", ".png", "gif":
			return "Images"
		case ".mov", ".mkv", ".mp4":
			return "Videos"
		case ".txt", ".doc", ".pdf":
			return "Documents"
		default:
			return "Others"
	}
}

// Function  to create a dir for a category
func createDir(dirPath, category string) error {
	categoryPath := filepath.Join(dirPath, category)
	return os.MkdirAll(categoryPath, os.ModePerm)
}

func moveFile(filePath, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath))
	return os.Rename(filePath, destPath)
}