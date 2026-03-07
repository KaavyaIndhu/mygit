package objects

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HashFile(filename string) {

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