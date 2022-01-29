package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2d struct {
	x int
	y int
}

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup = sync.WaitGroup{}
)

const noOfThreads int = 18

func findArea(inputChannel chan string) {
	for pointStr := range inputChannel {
		var points []Point2d
		regexData := r.FindAllStringSubmatch(pointStr, -1)

		for _, p := range regexData {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2d{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			// % is used so that we do not cross the limit of loop by doing i+1
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitGroup.Done()
}

func main() {
	absPath, _ := filepath.Abs("./")
	dat, _ := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(dat)

	inputChannel := make(chan string, 1000)
	for i := 0; i < noOfThreads; i++ {
		go findArea(inputChannel)
	}
	waitGroup.Add(noOfThreads)

	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		// split the text string file whenever a new line appears
		// findArea(line)
		inputChannel <- line
	}
	close(inputChannel) // signal to worker threads that work is finished
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("time taken %s", elapsed)
}

// string with 5 points of irregular polygon, whose are we will calculate using shoelace algo
// line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
