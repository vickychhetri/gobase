package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// main function is the entry point of the program.
func main() {

	// The folder path for which we want to calculate the total size.
	// NOTE: This includes all files and all subfolders recursively.
	folder := "C:\\Users\\ASUS\\Downloads\\"

	// Call GetSize() to calculate the total size.
	// It returns:
	//   - total size in bytes (int64)
	//   - an error (if any)
	total, err := GetSize(folder)
	if err != nil {
		// Panic stops the execution immediately and prints the error.
		panic(err)
	}

	// Print the final result to the console.
	fmt.Println("Total : ", fmt.Sprintf("%d", total))
}

// GetSize() calculates total size of a folder or files inside it.
// PARAMETERS:
//   - path: folder directory path
//
// RETURNS:
//   - total size in bytes (int64)
//   - error (if something goes wrong during reading the directory)
func GetSize(path string) (int64, error) {

	// os.ReadDir reads all items (files + folders) inside the given directory.
	// It returns a slice of DirEntry.
	contents, err := os.ReadDir(path)
	if err != nil {
		// Instead of returning error, the code panics here.
		// Ideally, return the error so the caller can handle it.
		panic(err)
	}

	// This variable will accumulate the total size of all files and subfolders.
	var total int64

	// Loop through every item inside the folder.
	for _, entry := range contents {

		// If the entry is a folder, recursively call GetSize() for that folder.
		// entry.IsDir() returns TRUE for directories.
		if entry.IsDir() {

			// Build the full path of the subfolder
			// filepath.Join handles correct slash formatting.
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}

			// Add the returned size of the subfolder to total.
			total += temp
		}

		// This block executes for ALL entries (files + directories).
		// It gets the basic metadata of the file/directory.
		info, err := entry.Info()
		if err != nil {
			return -1, err
		}

		// Add the size of the file to total.
		// For folders, info.Size() is normally zero on Windows.
		total += info.Size()
	}

	// Finally return the total calculated size.
	return total, nil
}
