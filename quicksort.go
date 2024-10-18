package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)

		quicksort(arr, pi+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high
	for i != j {
		for j != i && arr[j] >= pivot {
			j--
		}
		if j != i {
			arr[i], arr[j] = arr[j], arr[i]
		}
		for i != j && arr[i] <= pivot {
			i++
		}
		if j != i {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return i
}
func main() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		arr = append(arr, num)
	}
	fmt.Println("原始数组:", arr)

	quicksort(arr, 0, len(arr)-1)

	fmt.Println("排序后数组:", arr)
}
