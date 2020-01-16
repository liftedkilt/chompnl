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
		fmt.Println("Please pass the name of the file to chomp.")

		os.Exit(1)
	}

	fn := os.Args[1]

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

	if nlchars > 0 {
		log.Fatalln(fn + ": Removed " + strconv.Itoa(nlchars) + " newline(s).")
	}
}
