package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var prev byte
	total := int64(0)
	// scan a character
	// scan the next one and see if a match
	// if it is add the digit to the score and pop the top character
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	stats, err := file.Stat()
	check(err)
	contents := make([]byte, stats.Size())

	_, err = file.Read(contents)
	check(err)
	fmt.Println(len(contents))
	first := contents[0]

	for i, tok := range contents {
		val, err := strconv.ParseInt(string(tok), 10, 64)
		check(err)
		if tok == prev {
			total += val
		}
		if i == len(contents)-1 {
			if tok == first {
				total += val
			}
		}
		prev = tok
	}
	fmt.Print(total)

}
