package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsStrArr := strings.Split(scanner.Text(), " ")
	var numsIntArr []int64
	for _, value := range numsStrArr {
		v, _ := strconv.ParseInt(value, 10, 64)
		numsIntArr = append(numsIntArr, v)
	}
	numsIntArr = bubbleSort(numsIntArr)
	fmt.Println(numsIntArr)
}

func bubbleSort(arr []int64) []int64 {
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
