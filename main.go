package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"time"
	"strings"
)

func initRepo() {

	err := os.Mkdir(".mygit", 0755)
	if err != nil {
		fmt.Println("Repository already exists or cannot create:", err)
		return
	}

	os.Mkdir(".mygit/objects", 0755)
	os.Mkdir(".mygit/commits", 0755)

	headFile, _ := os.Create(".mygit/HEAD")
	headFile.Close()
	indexFile, _ := os.Create(".mygit/index")
    indexFile.Close()

	fmt.Println("Initialized empty MyGit repository")
}

func hashFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	hasher := sha1.New()
	io.Copy(hasher, file)

	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	objectPath := ".mygit/objects/" + hash

data, err := os.ReadFile(filename)
if err != nil {
	fmt.Println("Error reading file:", err)
	return
}

err = os.WriteFile(objectPath, data, 0644)
if err != nil {
	fmt.Println("Error writing object:", err)
	return
}
indexFile, err := os.OpenFile(".mygit/index", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
	fmt.Println("Error opening index:", err)
	return
}
defer indexFile.Close()

entry := filename + " " + hash + "\n"
indexFile.WriteString(entry)
fmt.Println("Stored object:", hash)
}
func commit(message string) {

	indexData, err := os.ReadFile(".mygit/index")
	if err != nil {
		fmt.Println("Error reading index:", err)
		return
	}

	if len(indexData) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	timestamp := time.Now().Unix()

	commitContent := "message: " + message + "\n"
	commitContent += "time: " + time.Now().Format(time.RFC3339) + "\n"
	commitContent += "files:\n"
	commitContent += string(indexData)

	commitFile := fmt.Sprintf(".mygit/commits/commit_%d", timestamp)

	err = os.WriteFile(commitFile, []byte(commitContent), 0644)
	if err != nil {
		fmt.Println("Error writing commit:", err)
		return
	}

	os.WriteFile(".mygit/HEAD", []byte(commitFile), 0644)

	os.WriteFile(".mygit/index", []byte(""), 0644)

	fmt.Println("Committed successfully:", commitFile)
}
func logCommits() {

	files, err := os.ReadDir(".mygit/commits")
	if err != nil {
		fmt.Println("Error reading commits:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No commits yet")
		return
	}

	for _, file := range files {

		path := ".mygit/commits/" + file.Name()

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading commit:", err)
			continue
		}

		fmt.Println("------")
		fmt.Println("commit:", file.Name())
		fmt.Println(string(data))
	}
}
func checkout(commitName string) {

	path := ".mygit/commits/" + commitName

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Commit not found:", err)
		return
	}

	lines := string(data)
	rows := strings.Split(lines, "\n")

	for _, row := range rows {

		if strings.Contains(row, ".txt") {

			parts := strings.Split(row, " ")
			filename := parts[0]
			hash := parts[1]

			objectPath := ".mygit/objects/" + hash

			fileData, err := os.ReadFile(objectPath)
			if err != nil {
				fmt.Println("Error reading object:", err)
				continue
			}

			os.WriteFile(filename, fileData, 0644)
			fmt.Println("Restored:", filename)
		}
	}
}
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: mygit <command>")
		return
	}

	command := os.Args[1]

	if command == "init" {
		initRepo()

	} else if command == "add" {

		if len(os.Args) < 3 {
			fmt.Println("Usage: mygit add <filename>")
			return
		}

		hashFile(os.Args[2])

	} else if command == "commit" {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit commit <message>")
		return
	}

	commit(os.Args[2])
	} else if command == "log" {

	logCommits()

} else if command == "checkout" {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit checkout <commit>")
		return
	}

	checkout(os.Args[2])
} else {
	fmt.Println("Unknown command:", command)
}
}