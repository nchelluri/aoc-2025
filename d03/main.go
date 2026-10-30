package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	partA := 0
	partB := 0

	for s.Scan() {
		l := s.Text()

		bs := strings.Split(l, "")

		bInts := make([]uint, 0, len(bs))

		for _, b := range bs {
			bInt, err := strconv.Atoi(b)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			bInts = append(bInts, uint(bInt))
		}

		// part B
		const numBatteries = 12

		maxIdxs := [numBatteries]int{}

		for i := range maxIdxs {
			maxIdxs[i] = -1
		}

		for i := range maxIdxs {
			firstIdx := 0
			lastIdx := len(bs) - numBatteries + i + 1

			for j := range i {
				firstIdx += maxIdxs[j] + 1
			}

			for j, b := range bInts[firstIdx:lastIdx] {
				if maxIdxs[i] == -1 || bInts[firstIdx:lastIdx][maxIdxs[i]] < b {
					maxIdxs[i] = j
				}
			}
		}

		bmjsB := ""

		for i := range numBatteries {
			idx := maxIdxs[i]

			for j := range i {
				idx += maxIdxs[j] + 1
			}

			bmjsB += fmt.Sprintf("%d", bInts[idx])
		}

		bmjB, err := strconv.Atoi(bmjsB)
		if err != nil {
			log.Fatalf("Error converting bank max joltage string (part B) to int: %v", err)
		}

		// part A
		firstMaxIdx, secondMaxIdx := -1, -1
		for i, b := range bInts[:len(bInts)-1] {
			if firstMaxIdx == -1 || bInts[:len(bInts)-1][firstMaxIdx] < b {
				firstMaxIdx = i
			}
		}

		for i, b := range bInts[firstMaxIdx+1:] {
			if secondMaxIdx == -1 || bInts[firstMaxIdx+1:][secondMaxIdx] < b {
				secondMaxIdx = i
			}
		}

		bmjsA := fmt.Sprintf("%d%d", bInts[firstMaxIdx], bInts[firstMaxIdx+1+secondMaxIdx])

		bmjA, err := strconv.Atoi(bmjsA)
		if err != nil {
			log.Fatalf("Error converting bank max joltage string to int: %v", err)
		}

		partA += bmjA
		partB += bmjB
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Println("part A", partA)
	fmt.Println("part B", partB)
}
