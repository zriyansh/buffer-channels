package main

import "fmt"

const matrixSize = 3

var (
	matrixA = [matrixSize][matrixSize]int{
		{3, 1, -4},
		{2, -3, 1},
		{5, -2, 0},
	}
	matrixB = [matrixSize][matrixSize]int{
		{3, 1, -4},
		{2, -3, 1},
		{5, -2, 0},
	}
	result = [matrixSize][matrixSize]int{}
)

func main() {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		fmt.Println(result[row])
	}
}
