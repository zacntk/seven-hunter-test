package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	decodeData()
	// encryptData()
}

func decodeData() {
	var encryptedData string
	fmt.Print("input = ")
	fmt.Scanln(&encryptedData)

	arr := strings.Split(encryptedData, "")

	decode := make([]int, len(arr)+1)

	unEncrypt := true

	for unEncrypt {
		for index := 0; index <= len(arr)-1; index++ {
			if arr[index] == "L" {
				if decode[index] > decode[index+1] {
					unEncrypt = false
				} else {
					decode[index] = decode[index+1] + 1
					unEncrypt = true
					break
				}
			} else if arr[index] == "R" {
				if decode[index] < decode[index+1] {
					unEncrypt = false
				} else {
					decode[index+1] += 1
					unEncrypt = true
					break
				}
			} else {
				if decode[index] == decode[index+1] {
					unEncrypt = false
				} else {
					if decode[index] > decode[index+1] {
						decode[index+1] = decode[index]
					} else {
						decode[index] = decode[index+1]
					}
					unEncrypt = true
					break
				}
			}
		}
	}

	ans := ""

	for _, value := range decode {
		ans += strconv.Itoa(value)
	}

	fmt.Println("output =", ans)
}

func encryptData() {
	var data string
	fmt.Print("input = ")
	fmt.Scanln(&data)

	arr := strings.Split(data, "")
	encrypt := make([]string, len(arr)-1)

	for index := 0; index <= len(arr)-2; index++ {
		left, _ := strconv.Atoi(arr[index])
		right, _ := strconv.Atoi(arr[index+1])
		if left > right {
			encrypt[index] = "L"
		} else if left < right {
			encrypt[index] = "R"
		} else {
			encrypt[index] = "="
		}
	}

	ans := ""
	for _, value := range encrypt {
		ans += value
	}

	fmt.Println("output =", ans)
}
