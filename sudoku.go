package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type sudoku struct {
	grid [9][9]Cell
}

// parse the sudoku file
func NewSudoku(filename string) *sudoku {
	s := &sudoku{}

	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lineScanner := bufio.NewScanner(buf)
	for i := 0; i < 9; i++ {
		lineScanner.Scan()
		row := lineScanner.Text()
		columns := strings.Split(row, ",")
		for j, c := range columns {
			var isFixed = true
			number, err := strconv.ParseUint(c, 10, 8)
			if err != nil {
				if c == "" {
					isFixed = false
					number = 0
				} else {
					log.Fatal(err)
				}
			}
			s.grid[i][j] = Cell{
				isFixed: isFixed,
				number:  uint8(number),
			}
		}
	}
	err = lineScanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (s sudoku) String() string {
	var sb strings.Builder

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.grid[i][j].number == 0 {
				sb.WriteString("   ")
			} else {
				sb.WriteString(" ")
				sb.WriteString(strconv.Itoa(int(s.grid[i][j].number)))
				sb.WriteString(" ")
			}
			if j%3 == 2 && j != 8 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
		if i%3 == 2 && i != 8 {
			sb.WriteString("-----------------------------\n")
		}
	}
	return sb.String()
}
