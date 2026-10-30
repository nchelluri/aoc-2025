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

	for s.Scan() {
		l := s.Text()

		seqs := strings.SplitSeq(l, ",")

		sumOfInvalids := 0

		for seq := range seqs {
			rng := strings.Split(seq, "-")

			start, end := rng[0], rng[1]

			startI, err := strconv.Atoi(start)
			if err != nil {
				log.Fatalf("Error converting start string to int: %v", err)
			}

			endI, err := strconv.Atoi(end)
			if err != nil {
				log.Fatalf("Error converting end string to int: %v", err)
			}

			for i := startI; i <= endI; i++ {
				candidate := fmt.Sprintf("%d", i)

				// part 1
				// if len(candidate)%2 != 0 {
				// 	continue
				// }

				// if candidate[:len(candidate)/2] == candidate[len(candidate)/2:] {
				// 	sumOfInvalids += i
				// }

				// part 2
				matches := make(map[string]bool)

				for j := 1; j <= len(candidate)/2; j++ {
					possibility := strings.Repeat(candidate[:j], len(candidate)/len(candidate[:j]))

					if matches[possibility] {
						continue
					}

					if candidate == possibility {
						matches[possibility] = true

						sumOfInvalids += i
					}
				}
			}
		}

		fmt.Println(sumOfInvalids)
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}
}
