package main

import (
	"fmt"
	"time"
)

type File []bool

type Chessboard map[string]File

func main() {
	nowTime := time.Now()
	fmt.Printf("%s\n", nowTime.Weekday())
	fmt.Printf("%s\n", nowTime.Month())
	fmt.Printf("%d\n", nowTime.Day())
	fmt.Printf("%d\n", nowTime.Year())

	hour, _, _ := nowTime.Clock()
	_, minute, _ := nowTime.Clock()

	fmt.Printf("%d\n", hour)
	fmt.Printf("%d\n", minute)
	//fmt.Print(time.Now())
	createBoaord()

}
func createBoaord() {
	newBoard := make(map[string][]bool)
	fileA := make([]bool, 4)
	fileB := make([]bool, 4)
	fileA[0] = true
	fileA[1] = false
	fileA[2] = false
	fileA[3] = true
	fileB[0] = true
	fileB[1] = true
	fileB[2] = true
	fileB[3] = true
	newBoard["A"] = fileA
	newBoard["B"] = fileB

	count := 0
	for _, v := range newBoard {
		if v[3] == true {
			count++
		}
	}
	fmt.Println(count)

}
func testFormatting(cb Chessboard) {
	count := 0
	for _, v := range cb {
		for _, i := range v {
			if i == true {
				count++
			}
		}

	}
}
