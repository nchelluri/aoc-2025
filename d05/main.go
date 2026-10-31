package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	start int
	end   int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	firstSection := true

	goodIngredients := make(map[rng]bool)

	fresh := 0

	maxIngredientID := 0

	for s.Scan() {
		l := s.Text()

		if l == "" {
			firstSection = false
			continue
		}

		if firstSection {
			rngStrs := strings.Split(l, "-")

			start, err := strconv.Atoi(rngStrs[0])
			if err != nil {
				log.Fatalf("Error converting start of range to int: %v", err)
			}

			end, err := strconv.Atoi(rngStrs[1])
			if err != nil {
				log.Fatalf("Error converting end of range to int: %v", err)
			}

			for existing := range goodIngredients {
				// expand right and left
				if start <= existing.start && end >= existing.end {
					delete(goodIngredients, existing)
				}

				// expand left
				if start <= existing.start && end >= existing.start && existing.end >= end {
					delete(goodIngredients, existing)
					end = existing.end
				}

				// expand right
				if start >= existing.start && start <= existing.end && existing.end <= end {
					delete(goodIngredients, existing)
					start = existing.start
				}
			}

			goodIngredients[rng{start: start, end: end}] = true

			if end > maxIngredientID {
				maxIngredientID = end
			}

			continue
		}

		ingredient, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Error converting ingredient to int: %v", err)
		}

		for rng := range goodIngredients {
			if rng.start <= ingredient && rng.end >= ingredient {
				fresh++
				break
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Println(fresh)

	numFreshIngredientIDs := 0
	for rng := range goodIngredients {
		skip := false

		for rng2 := range goodIngredients {
			if rng == rng2 {
				continue
			}

			// skip ranges that are contained within other ranges
			if rng.start >= rng2.start && rng.end <= rng2.end {
				skip = true
			}
		}

		if skip {
			continue
		}

		numFreshIngredientIDs += rng.end - rng.start + 1
	}

	fmt.Println(numFreshIngredientIDs)
}
