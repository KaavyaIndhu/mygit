package branch

import (
	"fmt"
	"os"
	"strings"
)

func CreateBranch(name string) {

	headData, err := os.ReadFile(".mygit/HEAD")
	if err != nil {
		fmt.Println("Error reading HEAD")
		return
	}

	currentBranch := strings.TrimSpace(string(headData))

	commitData, err := os.ReadFile(".mygit/branches/" + currentBranch)
	if err != nil {
		fmt.Println("Error reading current branch")
		return
	}

	newBranchPath := ".mygit/branches/" + name

	err = os.WriteFile(newBranchPath, commitData, 0644)
	if err != nil {
		fmt.Println("Error creating branch")
		return
	}

	fmt.Println("Branch created:", name)
}