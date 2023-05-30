package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateAlias() string {
	file, err := os.Open("./pkg/utils/aliases.txt")
	if err != nil {
		log.Print(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	var selectedWords []string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(words))
		pick := words[randomIndex]
		selectedWords = append(selectedWords, pick)
	}
	items := strings.Join(selectedWords, ".")
	return strings.ToUpper(items)
}

func GenerateCvu() string {
	rand.Seed(time.Now().UnixNano())

	// generates a random number up to the int64 limit
	firstMax := 999999999999999999
	firstMin := 100000000000000000
	firstNum := rand.Intn(firstMax-firstMin) + firstMin

	// generates a new random number to complete the 22 digits
	secMax := 9999
	secMin := 1000
	secNum := rand.Intn(secMax-secMin) + secMin

	// concatenates the given numbers
	cvu := fmt.Sprint(firstNum) + fmt.Sprint(secNum)
	return cvu
}
