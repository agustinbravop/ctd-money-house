package utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateAlias() string {
	file, err := os.Open("./aliases.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	words := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	selectedWords := []string{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(words))
		pick := words[randomIndex]
		selectedWords = append(selectedWords, pick)
	}

	items := strings.Join(selectedWords, ".")

	// fmt.Println(strings.ToUpper(items))
	return strings.ToUpper(items)
}
