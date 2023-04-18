package main

import (
	"bufio"
	_ "bufio"
	_ "bytes"
	"io"
	_ "io"
	"os"
	_ "os"
	"reflect"
	"strings"
	_ "strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestPrompt(t *testing.T) {
	expected := "->"
	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	w.Close()
	os.Stdout = temp
	in, _ := io.ReadAll(r)
	out := string(in)
	if out != expected {
		t.Errorf("%s expected but got %s", expected, out)
	}
}

func TestIntro(t *testing.T) {
	expected := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n->"
	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	w.Close()
	os.Stdout = temp

	in, _ := io.ReadAll(r)
	out := string(in)

	if out != expected {
		t.Errorf("%s expected but got %s", expected, out)
	}
}

func TestCheckNumbers(t *testing.T) {
	testCases := []struct {
		input      string
		expected   string
		shouldPass bool
	}{
		{"q\n", "", false},
		{"a\n", "Please enter a whole number!", true},
		{"4\n", "4 is not a prime number", true},
		{"3\n", "3 is a prime number", true},
	}

	for _, tc := range testCases {
		scanner := bufio.NewScanner(strings.NewReader(tc.input))
		out, answer := checkNumbers(scanner)
		if out != tc.expected || answer != tc.shouldPass {
			t.Errorf("Unexpected result for input %q: out=%q, answer=%v", tc.input, out, answer)
		}
	}
}

func TestReadUserInput(t *testing.T) {
	testNumbers := "7\nq\n"
	scanner := bufio.NewReader(strings.NewReader(testNumbers))
	doneChan := make(chan bool)

	go readUserInput(scanner, doneChan)

	var output []string

	for res := range doneChan {
		if res {
			return
		}
		line, _, err := scanner.ReadLine()
		if err != io.EOF {
			return
		}
		output = append(output, string(line))
	}

	expected := []string{"7 is a prime number", "Goodbye."}
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}
