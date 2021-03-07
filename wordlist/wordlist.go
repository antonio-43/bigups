package wordlist

import (
	"bufio"
	"log"
	"os"
)

func LoadWords(path string) (allWords []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("[!] ERROR: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		allWords = append(allWords, sc.Text())
	}

	return
}

func LoadFromSTDIN() (targets []string) {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		targets = append(targets, sc.Text())
	}

	return
}
