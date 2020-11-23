package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func isValid(a []int) bool {
	if len(a) < 3 {
		log.Fatal(a) // Yes, pickfire, this is an assert()
	}

	sort.Ints(a) // lazy
	// log.Println(a)
	return a[0]+a[1] > a[2]
}

func main() {
	log.Println(isValid([]int{5, 10, 25}) == false)

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a := make([]int, 3, 3)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Sscanf(l, "%d %d %d", &a[0], &a[1], &a[2])
		if isValid(a) {
			count++
		}
	}
	log.Println(count)

	// Part B
	file.Seek(0, 0) // Rewind
	scanner = bufio.NewScanner(file)

	b := make([][]int, 3, 3)
	for i := range b {
		b[i] = make([]int, 3, 3)
	}
	countb := 0
	for scanner.Scan() {
		i := 0
		l := scanner.Text()
		fmt.Sscanf(l, "%d %d %d", &(b[0][i]), &(b[1][i]), &(b[2][i]))
		// log.Println(b)

		scanner.Scan()
		l = scanner.Text()
		i++
		fmt.Sscanf(l, "%d %d %d", &(b[0][i]), &(b[1][i]), &(b[2][i]))

		scanner.Scan()
		l = scanner.Text()
		i++
		fmt.Sscanf(l, "%d %d %d", &(b[0][i]), &(b[1][i]), &(b[2][i]))

		// log.Println(b[0])
		if isValid(b[0]) {
			countb++
		}
		if isValid(b[1]) {
			countb++
		}
		if isValid(b[2]) {
			countb++
		}
	}
	log.Println(countb)

}
