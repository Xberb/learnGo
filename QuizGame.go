package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fileFlag := flag.String("file", "quiz_key.csv", "Path to quiz CSV file. Default: quiz_key.csv")
	timerFlag := flag.Int("timer", 30, "Time to complet quiz in seconds. Default is 30 seconds")
	flag.Parse()
	input, err := ioutil.ReadFile(*fileFlag)
	question, correct := 0, 0

	go trackTime(*timerFlag, &correct, *fileFlag)

	if err != nil {
		log.Fatal(err)
	}

	dataSet := csv.NewReader(strings.NewReader(string(input)))

	for {
		record, err := dataSet.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question %d %s: ", question+1, record[0])
		if getInput() == record[1] {
			correct += 1
		}
		question += 1
	}

	theEnd(correct, *fileFlag)

}
func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func trackTime(counter int, correct *int, file string) {
	timer := time.NewTimer(time.Duration(counter) * time.Second)
	<-timer.C
	theEnd(*correct, file)
}
func theEnd(correct int, file string) {
	question := getTotal(file)
	fmt.Printf("\nYou got %d out of %d questions correct!\n", correct, question)
	os.Exit(0)

}
func getTotal(file string) int {
	openfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	filedata, err := csv.NewReader(openfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return len(filedata)
}
