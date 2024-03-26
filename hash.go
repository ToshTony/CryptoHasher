package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func calculateSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)


	fmt.Println("Welcome to CryptoHasher - The SHA-256 Hash Calculator!")
	fmt.Print("Enter the string to calculate SHA-256 hash: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	sha256Hash := calculateSHA256(input)

	fmt.Println("SHA-256 Hash:", sha256Hash)
	fmt.Println("Thank you! Bye bye!")
}
