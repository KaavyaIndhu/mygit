package commits

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
	"time"
)
func CreateCommit(message string) {

	indexData, err := os.ReadFile(".mygit/index")
	if err != nil {
		fmt.Println("Error reading index:", err)
		return
	}

	if len(indexData) == 0 {
		fmt.Println("Nothing to commit")
		return
	}

	// read current branch
	branchData, err := os.ReadFile(".mygit/HEAD")
	if err != nil {
		fmt.Println("Error reading HEAD:", err)
		return
	}

	branch := strings.TrimSpace(string(branchData))

	// read parent commit from branch
	parent := ""
	commitData, err := os.ReadFile(".mygit/branches/" + branch)
	if err == nil {
		parent = strings.TrimSpace(string(commitData))
	}

	// build commit content
	commitContent := "message: " + message + "\n"
	commitContent += "time: " + time.Now().Format(time.RFC3339) + "\n"
	commitContent += "parent: " + parent + "\n"
	commitContent += "files:\n"
	commitContent += string(indexData)

	// generate commit hash
	hasher := sha1.New()
	hasher.Write([]byte(commitContent))
	commitHash := fmt.Sprintf("%x", hasher.Sum(nil))

	commitFile := ".mygit/commits/" + commitHash

	err = os.WriteFile(commitFile, []byte(commitContent), 0644)
	if err != nil {
		fmt.Println("Error writing commit:", err)
		return
	}

	// update branch pointer
	branchPath := ".mygit/branches/" + branch
	os.WriteFile(branchPath, []byte(commitHash), 0644)

	// clear staging area
	os.WriteFile(".mygit/index", []byte(""), 0644)

	fmt.Println("Committed successfully:", commitFile)
}