package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 300

var (
	matrixA   = [matrixSize][matrixSize]int{}
	matrixB   = [matrixSize][matrixSize]int{}
	result    = [matrixSize][matrixSize]int{}
	rwLock    = sync.RWMutex{}
	cond      = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	rwLock.RLock()
	for {
		waitGroup.Done()
		cond.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}
}

func main() {

	waitGroup.Add(matrixSize)
	// matrixSize = how many threads we will have
	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
		// fmt.Println(result[row])
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		waitGroup.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		waitGroup.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()

	}
	elaspsed := time.Since(start)
	fmt.Println(elaspsed)
}
