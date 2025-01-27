package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func findMaxPathFromBottom(arr [][]int) int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < len(arr[i-1]); j++ {
			if j < len(arr[i])-1 {
				arr[i-1][j] += max(arr[i][j], arr[i][j+1])
			} else {
				arr[i-1][j] += arr[i][j]
			}
		}
	}

	return arr[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSum() {
	input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	maxSum := findMaxPathFromBottom(input)
	fmt.Printf("The maximum sum from test#1 is: %d\n", maxSum)

	file, err := os.ReadFile("./files/hard.json")
	if err != nil {
		panic(err)
	}
	var input2 [][]int
	if err := json.Unmarshal(file, &input2); err != nil {
		panic(err)
	}
	maxSum2 := findMaxPathFromBottom(input2)
	fmt.Printf("The maximum sum from test#2 is: %d\n", maxSum2)

}
