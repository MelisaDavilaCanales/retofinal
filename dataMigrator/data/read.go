package data

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func ReadFile(personStringCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(personStringCh)

	file, err := os.Open("./data/source.data")
	if err != nil {
		log.Fatal("error al leer el archivo", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {

		// Skip the first line
		if lineNumber == 0 {
			lineNumber++
			continue
		}

		personStringCh <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
