package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var valStrs [][]string
	var ops []string

	for s.Scan() {
		l := s.Text()

		lineValStrs := strings.SplitSeq(l, "")

		var lineVals []string

		for valStr := range lineValStrs {
			if valStr == "*" || valStr == "+" {
				break
			}

			lineVals = append(lineVals, valStr)
		}

		if len(lineVals) > 0 {
			valStrs = append(valStrs, lineVals)
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

	var vals [][]int
	var colVals []int

	for i := len(valStrs[0]) - 1; i >= 0; i-- {
		valStr := ""

		for j := len(valStrs) - 1; j >= 0; j-- {
			if valStrs[j][i] != " " {
				valStr = valStrs[j][i] + valStr
			}
		}

		if valStr == "" {
			vals = append(vals, colVals)
			colVals = nil
		} else {
			val, err := strconv.Atoi(valStr)
			if err != nil {
				log.Fatalf("Error converting value string to int: %v", err)
			}

			colVals = append(colVals, val)
		}
	}

	slices.Reverse(ops)

	vals = append(vals, colVals)

	total := 0

	for i := 0; i < len(vals); i++ {
		subTotal := 1

		for j := 0; j < len(vals[i]); j++ {
			switch ops[i] {
			case "+":
				subTotal += vals[i][j]
			case "*":
				subTotal *= vals[i][j]
			}
		}

		if ops[i] == "+" {
			subTotal--
		}

		total += subTotal
	}

	fmt.Println(total)
}
