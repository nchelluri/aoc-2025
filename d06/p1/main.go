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

	var vals [][]int
	var ops []string

	for s.Scan() {
		l := s.Text()

		lineValStrs := strings.SplitSeq(l, " ")

		var lineVals []int

		for valStr := range lineValStrs {
			if valStr == "" {
				continue
			}

			if valStr == "*" || valStr == "+" {
				break
			}

			val, err := strconv.Atoi(valStr)
			if err != nil {
				log.Fatalf("Error converting value string to int: %v", err)
			}

			lineVals = append(lineVals, val)
		}

		if len(lineVals) > 0 {
			vals = append(vals, lineVals)
		} else {
			for sym := range strings.SplitSeq(l, " ") {
				if sym == "" {
					continue
				}

				ops = append(ops, sym)
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	total := 0
	for i := 0; i < len(vals[0]); i++ {
		subTotal := 1

		for j := 0; j < len(vals); j++ {
			switch ops[i] {
			case "*":
				subTotal *= vals[j][i]
			case "+":
				subTotal += vals[j][i]
			default:
				log.Fatalf("Unexpected operator: %s", ops[i])
			}
		}

		if ops[i] == "+" {
			subTotal--
		}

		total += subTotal
	}

	fmt.Println(total)
}
