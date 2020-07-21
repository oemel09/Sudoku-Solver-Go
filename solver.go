package main

import "errors"

type sudokuSolver struct {
	sudoku           *sudoku
	currentCellIndex int8
}

func NewSudokuSolver(filename string) *sudokuSolver {
	return &sudokuSolver{
		sudoku: NewSudoku(filename),
	}
}

func (s *sudokuSolver) solveAll() []sudoku {
	solutions := make([]sudoku, 0)
	for {
		solution, err := s.solve()
		if err != nil {
			break
		}
		solutions = append(solutions, solution)
		couldStepBack := s.stepBack()
		if !couldStepBack {
			break
		}
	}
	return solutions
}

func (s *sudokuSolver) solve() (sudoku, error) {
	for s.hasNext() {
		c := s.getCellByIndex(s.currentCellIndex)
		if c.isFixed {
			s.currentCellIndex++
		} else {
			isPossible := s.insertNextPossibleNumber(c)
			if !isPossible {
				couldStepBack := s.stepBack()
				if !couldStepBack {
					return sudoku{}, errors.New("no more possible solutions")
				}
			}
		}
	}
	return *s.sudoku, nil
}

func (s *sudokuSolver) hasNext() bool {
	return s.currentCellIndex < 81
}

func (s *sudokuSolver) insertNextPossibleNumber(currentCell *Cell) bool {
	for i := currentCell.oldNumber; i < 9; i++ {
		if s.canSetNumber(*currentCell, i+1) {
			currentCell.number = i + 1
			s.currentCellIndex++
			return true
		}
	}
	return false
}

func (s *sudokuSolver) canSetNumber(currentCell Cell, number uint8) bool {
	return currentCell.number == 0 && s.canSetInRow(number) && s.canSetInColumn(number) && s.canSetInSquare(number)
}

func (s *sudokuSolver) canSetInRow(number uint8) bool {
	row := s.currentCellIndex / 9
	canSet := true
	for j := 0; j < 9; j++ {
		if s.sudoku.grid[row][j].number == number {
			canSet = false
			break
		}
	}
	return canSet
}

func (s *sudokuSolver) canSetInColumn(number uint8) bool {
	column := s.currentCellIndex % 9
	canSet := true
	for i := 0; i < 9; i++ {
		if s.sudoku.grid[i][column].number == number {
			canSet = false
			break
		}
	}
	return canSet
}

func (s *sudokuSolver) canSetInSquare(number uint8) bool {
	row := (s.currentCellIndex / 27) * 3
	column := ((s.currentCellIndex / 3) % 3) * 3
	canSet := true

outer:
	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if s.sudoku.grid[i][j].number == number {
				canSet = false
				break outer
			}
		}
	}
	return canSet
}

func (s *sudokuSolver) stepBack() bool {
	isStepBack := s.stepBackToLastEditableCell()
	if !isStepBack {
		return false
	}
	s.getCellByIndex(s.currentCellIndex).resetNumber()
	s.resetAllNextCells()
	return true
}

func (s *sudokuSolver) stepBackToLastEditableCell() bool {
	ok := s.reduceCurrentCellIndex()
	if !ok {
		return false
	}
	for s.getCellByIndex(s.currentCellIndex).isFixed {
		ok = s.reduceCurrentCellIndex()
		if !ok {
			return false
		}
	}
	return true
}

func (s *sudokuSolver) reduceCurrentCellIndex() bool {
	s.currentCellIndex--
	if s.currentCellIndex < 0 {
		return false
	}
	return true
}

func (s *sudokuSolver) resetAllNextCells() {
	for i := s.currentCellIndex; i < 81-1; i++ {
		s.getCellByIndex(i + 1).resetOldNumber()
	}
}

func (s *sudokuSolver) getCellByIndex(index int8) *Cell {
	return &s.sudoku.grid[index/9][index%9]
}
