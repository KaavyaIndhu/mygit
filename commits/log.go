package commits

import (
	"fmt"
	"os"
	"strings"
)

func ShowLog() {

	branchData, err := os.ReadFile(".mygit/HEAD")
	if err != nil {
		fmt.Println("No commits yet")
		return
	}

	branch := strings.TrimSpace(string(branchData))

	commitData, err := os.ReadFile(".mygit/branches/" + branch)
	if err != nil {
		fmt.Println("No commits yet")
		return
	}

	current := strings.TrimSpace(string(commitData))

	for current != "" {

		path := ".mygit/commits/" + current

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading commit:", err)
			return
		}

		fmt.Println("------")
		fmt.Println("commit:", current)
		fmt.Println(string(data))

		lines := strings.Split(string(data), "\n")

		parent := ""

		for _, line := range lines {
			if strings.HasPrefix(line, "parent:") {
				parent = strings.TrimSpace(strings.TrimPrefix(line, "parent:"))
			}
		}

		parentPath := ".mygit/commits/" + parent

if _, err := os.Stat(parentPath); err != nil {
    break
}

current = parent
	}
}