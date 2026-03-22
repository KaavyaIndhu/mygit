package main

import (
    "fmt"
    "os"
    "strings"

    "mygit/repository"
    "mygit/objects"
	"mygit/commits"
	"mygit/branch"
)

func initRepo() {

	err := os.Mkdir(".mygit", 0755)
	if err != nil {
		fmt.Println("Repository already exists or cannot create:", err)
		return
	}

	os.Mkdir(".mygit/objects", 0755)
	os.Mkdir(".mygit/commits", 0755)
	os.Mkdir(".mygit/branches", 0755)

	// create staging index
	indexFile, _ := os.Create(".mygit/index")
	indexFile.Close()

	// initialize main branch
	os.WriteFile(".mygit/branches/main", []byte(""), 0644)

	// set HEAD to main
	os.WriteFile(".mygit/HEAD", []byte("main"), 0644)

	fmt.Println("Initialized empty MyGit repository")
}

func checkout(target string) {

	branchPath := ".mygit/branches/" + target

	if _, err := os.Stat(branchPath); err == nil {

		os.WriteFile(".mygit/HEAD", []byte(target), 0644)

		commitData, _ := os.ReadFile(branchPath)
		commitHash := strings.TrimSpace(string(commitData))

		restoreCommit(commitHash)

		fmt.Println("Switched to branch:", target)
		return
	}

	restoreCommit(target)
}
func restoreCommit(commitName string) {

	path := ".mygit/commits/" + commitName

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Commit not found:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, row := range lines {

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
		repository.InitRepo()

	} else if command == "add" {

		if len(os.Args) < 3 {
			fmt.Println("Usage: mygit add <filename>")
			return
		}

		objects.HashFile(os.Args[2])

	} else if command == "commit" {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit commit <message>")
		return
	}

	debug := false

    if len(os.Args) > 3 && os.Args[3] == "--debug" {
	debug = true
    }

commits.CreateCommit(os.Args[2], debug)
	} else if command == "log" {

	commits.ShowLog()

} else if command == "checkout" {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit checkout <commit>")
		return
	}

	checkout(os.Args[2])
	} else if command == "branch" {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit branch <name>")
		return
	}

	branch.CreateBranch(os.Args[2])

} else {
	fmt.Println("Unknown command:", command)
}
}