package turing

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestPalindrome(t *testing.T) {
	file, err := os.Open("../../palindrome.code")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(file)
	machine, err := NewFromReader(scanner)
	if err != nil {
		log.Fatalln(err)
		return
	}

	state, output := machine.Execute([]byte{'1', '0', '1'})
	outputStr := strings.ReplaceAll(string(output), "_", "")
	if state != AcceptStatus || outputStr != "" {
		t.Error("Error happens")
	}

	machine.Reset()

	state, output = machine.Execute([]byte{'1', '0', '0'})
	outputStr = strings.ReplaceAll(string(output), "_", "")
	if state != ErrorStatus || outputStr != "00" {
		t.Error("Error happenes")
	}
}

func TestWriteTripleOne(t *testing.T) {
	file, err := os.Open("../../ones.code")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(file)
	machine, err := NewFromReader(scanner)
	if err != nil {
		log.Fatalln(err)
		return
	}

	state, output := machine.Execute([]byte{})
	outputStr := strings.ReplaceAll(string(output), "_", "")
	if state != AcceptStatus || outputStr != "111" {
		t.Error("Error happens")
	}
}
