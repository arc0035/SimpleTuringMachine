package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"turingmachine/src/turing"
)

func main() {
	file, err := os.Open("palindrome.code")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(file)
	machine, err := turing.NewFromReader(scanner)
	if err != nil {
		log.Fatalln(err)
		return
	}

	state, output := machine.Execute([]byte{'1', '0', '1'})
	fmt.Printf("state %v \noutput %v\n", state, strings.ReplaceAll(string(output), "_", ""))
	// optionally, resize scanner's capacity for lines over 64K, see next example
}
