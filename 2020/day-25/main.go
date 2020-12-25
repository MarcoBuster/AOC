package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bruteforceHandshake(subjectNumber int, publicKey int) int {
	value := 1
	for i := 1; true; i++ {
		value *= subjectNumber
		value %= 20201227
		if value == publicKey {
			return i
		}
	}
	return -1
}

func transformSubjectNumber(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value %= 20201227
	}
	return value
}

func getEncryptionKey(cardPK, cardLS, doorPK, doorLS int) (int, int) {
	cardEK := transformSubjectNumber(doorPK, cardLS)
	doorEK := transformSubjectNumber(cardPK, doorLS)
	return cardEK, doorEK
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	b.Scan()
	cardPK, _ := strconv.Atoi(b.Text())
	b.Scan()
	doorPK, _ := strconv.Atoi(b.Text())

	cardLS := bruteforceHandshake(7, cardPK)
	doorLS := bruteforceHandshake(7, doorPK)
	ek, _ := getEncryptionKey(cardPK, cardLS, doorPK, doorLS)
	fmt.Println("Result (part 1):", ek)
}
