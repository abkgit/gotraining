package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body) // Create a new scanner that takes data in Body as input.
	sc.Split(bufio.ScanWords)        // Split the scanned material acc. to space-separated words.

	for sc.Scan() { // Iterate by words as set above and put them into our map.
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil { // If there was an error other than EOF, report here.
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words { // Prints first 200 words from the map.
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
