package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	easy := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	findMaxPath(easy)

	file, err := os.Open("hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var hard [][]int
	err = json.Unmarshal(byteValue, &hard)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	findMaxPath(hard)
}

func findMaxPath(arr [][]int) {
	fmt.Println("input = ", arr)
	rows := len(arr)
	for row := rows - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			underLeft := arr[row+1][col]
			underRight := arr[row+1][col+1]
			arr[row][col] += max(underLeft, underRight) //Top
		}
	}
	fmt.Println("output =", arr[0][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
