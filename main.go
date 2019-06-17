package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	//version1()
	//version2()
	//version3()
	//version4()
	version5()
}

const (
	defaultTimeout = 30
)

// Ask questions, no timer
func version1() {
	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	var questionsAsked int
	var score int

	for _, line := range lines {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("What is %s?\n", line[0])
		questionsAsked++
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")

		if answer == line[1] {
			score++
		}
	}

	fmt.Printf("You got %d out of %d\n", score, questionsAsked)
}

// Ask questions, timer, but doesn't time out when waiting for an answer
func version2() {
	// read in csv
	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	var questionsAnswered int

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Hello, press Enter to start timer\n")
	s.Scan() // they pressed enter
	start := time.Now()

	timeout := time.After(10 * time.Second)
	//tick := time.Tick(500 * time.Millisecond)
	var score *int
	zero := 0
	score = &zero

loop:
	for {
		select {
		case <-timeout:
			// run out of time
			break loop
		default:
			if questionsAnswered == len(lines) {
				break
			}

			askQuestion(lines[questionsAnswered], score)

			questionsAnswered++
		}

	}

	now := time.Now()
	elapsed := now.Sub(start).Seconds()
	//t:=timer.Stop()
	fmt.Printf("You took %f seconds.\n", elapsed)
	fmt.Printf("You got %d out of %d.\n", *score, questionsAnswered)
}

// Read flags, ask questions, timeout works
func version3() {
	// TODO Add shuffling of questions https://yourbasic.org/golang/shuffle-slice-array/
	timeoutPtr := flag.Int("t", 30, "number of seconds until quiz times out")

	flag.Parse()

	timeout := defaultTimeout
	if timeoutPtr != nil {
		timeout = *timeoutPtr
	}

	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("You will have %d second(s) to complete the quiz. Press Enter.", timeout)
	s.Scan()
	timeUp := time.After(time.Duration(timeout) * time.Second)

	var questionsAsked int
	var score int

	done := make(chan bool)
	go func(done chan bool) {
		for _, line := range lines {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("What is %s?\n", line[0])
			questionsAsked++
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSuffix(answer, "\n")

			if answer == line[1] {
				score++
			}
		}

		done <- true
	}(done)

	select {
	case <-timeUp:
		// run out of time
		fmt.Print("\nOut of time\n")
		break
	case <-done:
		// all questions answered
		break
	}

	fmt.Printf("You got %d out of %d\n", score, questionsAsked)
}

// Read flags, ask questions, times out, shuffles using rand
func version4() {
	// TODO Add shuffling of questions https://yourbasic.org/golang/shuffle-slice-array/
	var timeoutPtr *int
	//timeoutPtr := flag.Int("t", 30, "number of seconds until quiz times out")
	//
	//flag.Parse()

	timeout := defaultTimeout
	if timeoutPtr != nil {
		timeout = *timeoutPtr
	}

	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	// Shuffle questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })

	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("You will have %d second(s) to complete the quiz. Press Enter.", timeout)
	s.Scan()
	timeUp := time.After(time.Duration(timeout) * time.Second)

	var questionsAsked int
	var score int

	done := make(chan bool)
	go func(done chan bool) {
		for _, line := range lines {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("What is %s?\n", line[0])
			questionsAsked++
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSuffix(answer, "\n")

			if answer == line[1] {
				score++
			}
		}

		done <- true
	}(done)

	select {
	case <-timeUp:
		// run out of time
		fmt.Print("\nOut of time\n")
		break
	case <-done:
		// all questions answered
		break
	}

	fmt.Printf("You got %d out of %d\n", score, questionsAsked)
}

// Read flags, ask questions, times out, shuffles using rand
func version5() {
	var timeoutPtr *int
	//timeoutPtr := flag.Int("t", 30, "number of seconds until quiz times out")
	//
	//flag.Parse()

	timeout := defaultTimeout
	if timeoutPtr != nil {
		timeout = *timeoutPtr
	}

	f, err := os.Open("problems.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	//Put questions into map so that we get a random order when iterating
	lineMap := make(map[int][]string)
	for i, line := range lines {
		lineMap[i] = line
	}

	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("You will have %d second(s) to complete the quiz. Press Enter.", timeout)
	s.Scan()
	timeUp := time.After(time.Duration(timeout) * time.Second)

	var questionsAsked int
	var score int

	done := make(chan bool)
	go func(done chan bool) {
		for _, line := range lineMap {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("What is %s?\n", line[0])
			questionsAsked++
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSuffix(answer, "\n")

			if answer == line[1] {
				score++
			}
		}

		done <- true
	}(done)

	select {
	case <-timeUp:
		// run out of time
		fmt.Print("\nOut of time\n")
		break
	case <-done:
		// all questions answered
		break
	}

	fmt.Printf("You got %d out of %d\n", score, questionsAsked)
}

func askQuestion(questionAndAnswer []string, score *int) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("What is %s?\n", questionAndAnswer[0])

	answer, _ := reader.ReadString('\n')

	answer = strings.TrimSuffix(answer, "\n")

	if answer == questionAndAnswer[1] {
		*score++
	}

}
