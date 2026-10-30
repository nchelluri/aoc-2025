package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	pos := 50
	numZeroPos := 0
	crossedZero := 0

	for s.Scan() {
		l := s.Text()

		factor := 1

		switch l[0] {
		case 'L':
			factor = -1
		case 'R':
			factor = 1
		default:
			log.Fatalf("Unexpected direction: %c", l[0])
		}

		moveBy, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatalf("Unexpected error for moveBy %s: %v", l[1:], err)
		}

		var cz int
		pos, cz = move(pos, factor, moveBy)

		crossedZero += cz

		if pos == 0 {
			numZeroPos++
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Error with scanner: %v", err)
	}

	fmt.Printf("numZeroPos:  %d\n", numZeroPos)
	fmt.Printf("crossedZero: %d\n", crossedZero)
}

func mod(num, mod int) int {
	if num < 0 {
		for num < 0 {
			num += mod
		}
	} else if num >= mod {
		for num >= mod {
			num -= mod
		}
	}

	return num
}

func move(pos, factor, moveBy int) (int, int) {
	delta := factor * moveBy

	crossedZero := 0

	for i := pos; i != pos+delta; i += factor {
		if mod(i, 100) == 0 {
			crossedZero++
		}
	}

	pos += delta

	pos = mod(pos, 100)

	return pos, crossedZero
}
