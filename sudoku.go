package main

import (
	"fmt"
)

var initialBoard = [9][9]int8{
	{3, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 0, 8, 0, 2, 0},
	{0, 0, 0, 0, 9, 5, 0, 0, 0},
	{0, 0, 0, 4, 0, 1, 0, 9, 7},
	{2, 0, 0, 8, 7, 3, 4, 0, 1},
	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 4},
	{0, 0, 5, 0, 4, 0, 3, 0, 0},
	{0, 0, 0, 0, 0, 2, 7, 1, 0},
}

type SudokuBoard struct {
	grid       [9][9]int8
	candidates [9][9][]int8
}

func (b *SudokuBoard) Init(startGrid [9][9]int8) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.grid[i][j] = startGrid[i][j]
		}
	}
	b.updateCandidates()
}

func (b *SudokuBoard) updateCandidates() {
	// produce a list of possible candidate values for each element of the grid
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.grid[i][j] > 0 {
				continue //skip the cells already assigned with a value
			}
			var val int8
			for val = 1; val < 10; val++ {
				if b.isValidAssignment(i, j, val) {
					b.candidates[i][j] = append(b.candidates[i][j], val)
				}
			}
		}
	}
}

func (b *SudokuBoard) isValidAssignment(row int, col int, val int8) bool {

	//check column for existing val instances, return false if found any
	for i := 0; i < 9; i++ {
		if b.grid[i][col] == val {
			return false
		}
	}
	//check row for existing val instances, return false if found any
	for j := 0; j < 9; j++ {
		if b.grid[row][j] == val {
			return false
		}
	}
	//check submatrix 3x3 corresponding row,col for val instance, return false if found any
	var firstRow, firstCol int
	firstRow = row / 3
	firstCol = col / 3
	for i := firstRow * 3; i < firstRow*3+3; i++ {
		for j := firstCol * 3; j < firstCol*3+3; j++ {
			if b.grid[i][j] == val {
				return false
			}
		}
	}
	return true

}

func (b *SudokuBoard) IsResolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *SudokuBoard) String() string {
	//override standard print formatting
	s := ""
	fmtStr := "%d "
	tmp := ""
	resolved := b.IsResolved()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (j+1)%3 == 0 {
				fmtStr = "%2d %-14s  "
			} else {
				fmtStr = "%2d %-14s "
			}
			if len(b.candidates[i][j]) > 0 && !resolved {
				tmp = fmt.Sprintf("%+v", b.candidates[i][j])
			} else {
				tmp = ""
			}

			s += fmt.Sprintf(fmtStr, b.grid[i][j], tmp)
		}
		s += "\n"
		if (i+1)%3 == 0 {
			s += "\n"
		}
	}
	return s
}

func (b *SudokuBoard) Resolve(i, j int) bool {

	if j > 8 {
		i++
		j = 0
	}
	if i > 8 {
		return true
	}
	if b.grid[i][j] > 0 {
		return b.Resolve(i, j+1)
	}
	//	for val :=1; val <=9 ; val++ {
	for _, val := range b.candidates[i][j] {
		if b.isValidAssignment(i, j, int8(val)) {
			b.grid[i][j] = int8(val)
			if b.Resolve(i, j+1) {
				return true
			}
			b.grid[i][j] = 0

		}
	}

	return false
}

func main() {
	sudoku := new(SudokuBoard)
	sudoku.Init(initialBoard)
	fmt.Println(sudoku)

	if sudoku.Resolve(0, 0) {
		fmt.Println("Found a solution")
		fmt.Println(sudoku)
	} else {
		fmt.Println("Did NOT find any solution")
		fmt.Println(sudoku)

	}

}
