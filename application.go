package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "csv/sudoku_03.csv"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	solver := NewSudokuSolver(filename)

	start := time.Now()
	solutions := solver.solveAll()
	elapsed := time.Since(start)

	for _, s := range solutions {
		fmt.Println(s)
	}
	fmt.Printf("Took %s to calculate %v solution(s)", elapsed, len(solutions))
}
