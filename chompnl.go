package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please pass the names of the file to chomp.")

		os.Exit(1)
	}

	filesChanged := 0

	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]

		nlchars := chompfile(fn)

		if nlchars > 0 {
			filesChanged++
			log.Println(fn + ": Removed " + strconv.Itoa(nlchars) + " newline(s).")
		}
	}

	if filesChanged > 0 {
		log.Fatalln("Removed newlines from " + strconv.Itoa(filesChanged) + " files.")
	}
}

func chompfile(fn string) int {
	b, err := ioutil.ReadFile(fn) // b has type []byte
	if err != nil {
		log.Fatal(err)
	}

	lb := len(b)

	nlchars := 0
	for i := 1; i < lb; i++ {
		if b[lb-i] == '\n' {
			nlchars++
		} else {
			break
		}
	}

	out := b[:lb-nlchars]

	ioutil.WriteFile(fn, out, 0644)

	return nlchars
}
