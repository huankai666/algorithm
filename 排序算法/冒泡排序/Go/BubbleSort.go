package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsArr := strings.Split(scanner.Text(), " ")
	fmt.Println(numsArr)
	//result := bubbleSort([]int{numsArr})
}

func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
