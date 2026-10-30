package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var grid [][]bool

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		l := s.Text()

		rowStrs := strings.Split(l, "")

		row := make([]bool, 0, len(rowStrs))

		for _, cell := range rowStrs {
			switch cell {
			case "@":
				row = append(row, true)
			case ".":
				row = append(row, false)
			default:
				log.Fatalf("Unexpected cell string: %s", cell)
			}
		}

		grid = append(grid, row)
	}

	// numAccessible := 0

	numRemovable := 0
	removable := true

	for removable {
		removable = false

		for i := range grid {
			for j := range grid[i] {
				if !grid[i][j] {
					continue
				}

				numAdj := 0

				// row above
				if i > 0 {
					if j > 0 {
						if grid[i-1][j-1] {
							numAdj++
						}
					}

					if grid[i-1][j] {
						numAdj++
					}

					if j < len(grid[i-1])-1 {
						if grid[i-1][j+1] {
							numAdj++
						}
					}
				}

				// same row
				if j > 0 {
					if grid[i][j-1] {
						numAdj++
					}
				}

				if j < len(grid[i])-1 {
					if grid[i][j+1] {
						numAdj++
					}
				}

				// row below
				if i < len(grid)-1 {
					if j > 0 {
						if grid[i+1][j-1] {
							numAdj++
						}
					}

					if grid[i+1][j] {
						numAdj++
					}

					if j < len(grid[i+1])-1 {
						if grid[i+1][j+1] {
							numAdj++
						}
					}
				}

				if numAdj < 4 {
					// numAccessible++

					numRemovable++
					grid[i][j] = false
					removable = true
				}
			}
		}

	}

	if err := s.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Println(numRemovable)
}
