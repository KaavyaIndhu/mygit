package repository

import (
	"fmt"
	"os"
)

func InitRepo() {

	err := os.Mkdir(".mygit", 0755)
	if err != nil {
		fmt.Println("Repository already exists or cannot create:", err)
		return
	}

	os.Mkdir(".mygit/objects", 0755)
	os.Mkdir(".mygit/commits", 0755)
	os.Mkdir(".mygit/branches", 0755)

	headFile, _ := os.Create(".mygit/HEAD")
	headFile.Close()

	indexFile, _ := os.Create(".mygit/index")
	indexFile.Close()

	os.WriteFile(".mygit/branches/main", []byte(""), 0644)
	os.WriteFile(".mygit/HEAD", []byte("main"), 0644)

	fmt.Println("Initialized empty MyGit repository")
}