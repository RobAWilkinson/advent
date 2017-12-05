package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("table.txt")
	defer file.Close()
	check(err)

	reader := bufio.NewReader(file)
	total := int64(0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		row := strings.Fields(string(line))
		highest := int64(0)
		lowest := int64(0)

		for i, strval := range row {
			val, _ := strconv.ParseInt(strval, 10, 64)
			if i == 0 {
				lowest = val
			}
			if val > highest {
				highest = val
			}
			if val < lowest {
				lowest = val
			}
		}
		diff := highest - lowest
		total += diff
	}
	fmt.Println(total)

}
