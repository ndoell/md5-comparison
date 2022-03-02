package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

	// Hash to compare.
	hash := "6a3e2e891abfa787c3ccfdedde8241ca"

	// Open text file to read.
	f, err := os.Open("gotclean.txt")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Append words from file to slice.
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Cycle through wordlist using every combination and compare to MD5.
	for _, str1 := range words {
		for _, str2 := range words {
			password := str1 + str2
			hashpw := GetMD5Hash(password)

			result := (strings.Compare(hash, hashpw))

			// If result = 0 strings matched.
			if result == 0 {
				fmt.Println("Answer: " + str1 + str2)
			}
		}
	}

}
